package database

import (
	"fmt"

	"github.com/yongchengchen/sysgo/contract"
	"github.com/yongchengchen/sysgo/services/container"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
	"github.com/mitchellh/mapstructure"
)

func connect(name string) *gorose.Engin {
	if conf, ok := container.Get("config").(contract.IConfig); ok {
		configs := conf.Get("database.connections." + name)
		var dbConf gorose.Config
		if err := mapstructure.Decode(configs, &dbConf); err != nil {
			fmt.Printf("database.connection.%#v is not correct %#v\n", name, configs)
		}
		if engin, err := gorose.Open(&dbConf); err == nil {
			return engin
		}
	}
	return nil
}

func Connection(name string) *gorose.Engin {
	key := "db_" + name
	if ins := container.Get(key); ins != nil {
		if conn, ok := container.Get(key).(*gorose.Engin); ok {
			return conn
		}
	}

	if conn := connect(name); conn != nil {
		container.Put(key, conn)
		return conn
	}
	return nil
}

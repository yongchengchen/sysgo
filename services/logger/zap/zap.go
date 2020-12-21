// package zap handles creating zap logger
package zap

import (
	"fmt"

	"github.com/yongchengchen/sysgo/contract"
	"github.com/yongchengchen/sysgo/services/container"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func RegisterLogDriver() (interface{}, error) {
	zLogger, err := initLog()
	if err != nil {
		return nil, errors.Wrap(err, "RegisterLogrusLog")
	}
	defer zLogger.Sync()
	zSugarlog := zLogger.Sugar()
	zSugarlog.Info()

	//This is for loggerWrapper implementation
	//appLogger.SetLogger(&loggerWrapper{zaplog})

	//SetLogger(zSugarlog)
	return zSugarlog, nil
}

func Factory() interface{} {
	if zSugarlog, err := RegisterLogDriver(); err == nil {
		return zSugarlog
	}
	return nil
}

// initLog create logger
func initLog() (*zap.Logger, error) {
	c := container.Get("config").(contract.IConfig)

	configItem := c.Get("logging.channels.zap")
	var cfg zap.Config
	if err := mapstructure.Decode(configItem, &cfg); err != nil {
		fmt.Printf("zapConfig!!!: %#v\n", cfg)
	}
	var zLogger *zap.Logger
	// //customize it from configuration file
	err := customizeLogFromConfig(&cfg, c)
	if err != nil {
		return zLogger, errors.Wrap(err, "cfg.Build()")
	}
	zLogger, err = cfg.Build()
	if err != nil {
		return zLogger, errors.Wrap(err, "cfg.Build()")
	}

	zLogger.Debug("logger construction succeeded")
	return zLogger, nil
}

// customizeLogrusLogFromConfig customize log based on parameters from configuration file
func customizeLogFromConfig(zapConf *zap.Config, appConf contract.IConfig) error {
	// set log level
	if level, ok := appConf.Get("logging.channels.zap.level").(string); ok {
		atomicLevel := zap.NewAtomicLevel()
		l := atomicLevel.Level()
		l.Set(level)
		atomicLevel.SetLevel(l)
		zapConf.Level = atomicLevel
	}
	return nil
}

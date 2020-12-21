package models

import (
	"time"

	"github.com/yongchengchen/sysgo/services/database"
	"github.com/gohouse/gorose/v2"
)

type UrlRewrite struct {
	ID        int       `gorose:"id"`
	CreatedAt time.Time `gorose:"created_at"`
	UpdatedAt time.Time `gorose:"updated_at"`

	StoreID      int    `gorose:"store_id"`
	RequestPath  string `gorose:"request_path"`
	Hash         string `gorose:"hash"`
	TargetType   string `gorose:"target_type"`
	TargetPath   string `gorose:"target_path"`
	RedirectType string `gorose:"redirect_type"`
	Description  string `gorose:"description"`
	Metadata     string `gorose:"metadata"`
}

// 设置表名, 如果没有设置, 默认使用struct的名字
func (m *UrlRewrite) TableName() string {
	return "ink_url_rewrites"
}

func (m *UrlRewrite) Connection() *gorose.Engin {
	return database.Connection("mysql")
}

func (m *UrlRewrite) Orm() gorose.IOrm {
	return database.Connection("mysql").NewOrm()
}

func (m *UrlRewrite) Session() gorose.ISession {
	return database.Connection("mysql").NewSession()
}

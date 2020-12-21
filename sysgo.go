package main

import (
	"github.com/yongchengchen/sysgo/app"
	"github.com/yongchengchen/sysgo/services/sysstat"
)

func main() {
	app.InitServices()
	sysstat.Post()
}

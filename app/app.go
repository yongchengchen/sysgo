package app

import (
	"github.com/yongchengchen/sysgo/contract"
	"github.com/yongchengchen/sysgo/services/config"
	"github.com/yongchengchen/sysgo/services/container"
	"github.com/yongchengchen/sysgo/services/logger"
	"github.com/yongchengchen/sysgo/services/user"
)

func InitServices() {
	container.Bind("config", config.Factory, true)
	container.Bind("log", logger.Factory, true)
	// container.Bind("router", router.Factory, true)
	container.Bind("user", user.Factory, true)
}

func ConfigService() contract.IConfig {
	return container.Get("config").(contract.IConfig)
}

func LogService() contract.ILogger {
	return container.Get("log").(contract.ILogger)
}

// func RouterService() *gin.Engine {
// 	return container.Get("router").(*gin.Engine)
// }

// func Run() {
// 	if services, ok := ConfigService().Get("app.boot_services").([]interface{}); ok {
// 		container.Boot(services)
// 	}
// 	RouterService().Run(":8080")
// }

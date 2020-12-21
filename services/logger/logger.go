package logger

import (
	"github.com/yongchengchen/sysgo/contract"
	"github.com/yongchengchen/sysgo/services/container"
	"github.com/yongchengchen/sysgo/services/logger/zap"
)

type Logger struct {
}

func Factory() interface{} {
	registerDrivers()
	l := Logger{}
	return &l
}

func registerDrivers() {
	container.Bind("zap_logger", zap.Factory, true)
	// container.Bind('logrus', Zap.Factory, true)
}

func (c *Logger) Info(args ...interface{}) {
	_loggor := container.Get("zap_logger")
	if log, ok := _loggor.(contract.ILogger); ok {
		log.Info(args)
	}
}

func (c *Logger) Infof(message string, args ...interface{}) {
	if _loggor, ok := container.Get("zap_logger").(contract.ILogger); ok {
		_loggor.Infof(message, args...)
	}
}

func (c *Logger) Errorf(message string, args ...interface{}) {
	if _loggor, ok := container.Get("zap_logger").(contract.ILogger); ok {
		_loggor.Errorf(message, args...)
	}
}

func (c *Logger) Fatal(args ...interface{}) {
	if logger, ok := container.Get("zap_logger").(contract.ILogger); ok {
		logger.Fatal(args)
	}
}

func (c *Logger) Fatalf(message string, args ...interface{}) {
	if _loggor, ok := container.Get("zap_logger").(contract.ILogger); ok {
		_loggor.Fatalf(message, args...)
	}
}

func (c *Logger) Warnf(message string, args ...interface{}) {
	if _loggor, ok := container.Get("zap_logger").(contract.ILogger); ok {
		_loggor.Warnf(message, args)
	}
}

func (c *Logger) Debug(args ...interface{}) {
	if _loggor, ok := container.Get("zap_logger").(contract.ILogger); ok {
		_loggor.Debug("Debug", args)
	}
}

func (c *Logger) Debugf(message string, args ...interface{}) {
	if logger, ok := container.Get("zap_logger").(contract.ILogger); ok {
		logger.Debug(message, args)
	}
}

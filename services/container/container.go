package container

import (
	"errors"
	"fmt"

	"github.com/yongchengchen/sysgo/contract"
)

var (
	providers          map[string]interface{}
	singletonProviders map[string]interface{}
	singletonResolved  map[string]interface{}
)

func init() {
	providers = make(map[string]interface{})
	singletonProviders = make(map[string]interface{})
	singletonResolved = make(map[string]interface{})
}

func Bind(name string, factory func() interface{}, singleton bool) error {
	if factory == nil {
		return errors.New("can't provide an untyped nil")
	}
	if _, ok := singletonResolved[name]; ok {
		return errors.New(fmt.Sprintf("provider '%s' has already been resolved. Make sure bind before it's been resolved.", name))
	}

	if singleton {
		singletonProviders[name] = factory
	} else {
		providers[name] = factory
	}
	return nil
}

func Get(name string) interface{} {
	if ins, ok := singletonResolved[name]; ok {
		return ins
	}
	if constructor, ok := providers[name]; ok {
		if f, ok := constructor.(func() interface{}); ok {
			return f()
		}
		return nil
	}

	if constructor, ok := singletonProviders[name]; ok {
		if f, ok := constructor.(func() interface{}); ok {
			ins := f()
			singletonResolved[name] = ins
			return ins
		}
	}
	return nil
}

func Put(name string, instance interface{}) {
	singletonResolved[name] = instance
}

func Boot(services []interface{}) {
	for _, key := range services {
		if name, ok := key.(string); ok {
			if service, ok := Get(name).(contract.IService); ok {
				service.Boot()
			}
		}
	}
}

package commons

import (
	"fmt"
	"reflect"
)

type PluginRegistry struct {
	registry map[string]reflect.Type
	count    int32
}

func NewPluginRegistry() *PluginRegistry {
	return &PluginRegistry{
		registry: make(map[string]reflect.Type),
		count:    0,
	}
}

func (r *PluginRegistry) GetPluginCount() int32 {
	return r.count
}

func (r *PluginRegistry) Register(name string, typ reflect.Type) {
	r.registry[name] = typ
	r.count++
}

func (r *PluginRegistry) CreatePlugin(name string, whitelist []string) (Plugin, error) {
	if typ, ok := r.registry[name]; ok {
		instance := reflect.New(typ).Interface()
		if plg, ok := instance.(Plugin); ok {
			plg.SetWhiteList(whitelist)
			return plg, nil
		} else {
			return nil, fmt.Errorf("插件%s错误", name)
		}
	}

	return nil, fmt.Errorf("插件%s未注册", name)
}

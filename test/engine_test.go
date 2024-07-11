package test

import (
	"qqbot-reconstruction/internal/app/plugins"
	"reflect"
	"testing"
)

func TestEngine(t *testing.T) {
	pluginRegistry := plugins.NewPluginRegistry()
	pluginRegistry.Register("music", reflect.TypeOf(plugins.MusicPlugin{}))
	msgEngine := plugins.NewPluginEngine()
	msgEngine.Init("E:\\Projects\\qqbot-reconstruction\\configs\\plugins.json", pluginRegistry)
}

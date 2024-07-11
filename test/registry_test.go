package test

import (
	"fmt"
	"qqbot-reconstruction/internal/app/plugins"
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	pluginRegistry := plugins.NewPluginRegistry()

	// 注册类型: 映射配置中的插件名与插件类
	pluginRegistry.Register("music", reflect.TypeOf(plugins.MusicPlugin{}))

	// 创建类型：运行时根据配置文件动态加载插件
	if imagePlugin, err := pluginRegistry.CreatePlugin("image"); err == nil {
		imagePlugin.GetWhiteList()
	}
	if musicPlugin, err := pluginRegistry.CreatePlugin("music"); err == nil {
		musicPlugin.GetWhiteList()
	}
}

func TestScanAndRegister(t *testing.T) {
	pluginRegistry := plugins.NewPluginRegistry()
	_ = plugins.ScanAndRegisterPlugins("E:\\Projects\\qqbot-reconstruction\\internal\\app\\plugins", pluginRegistry)
	fmt.Println(pluginRegistry.GetPluginCount())
}

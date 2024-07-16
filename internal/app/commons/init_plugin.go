package commons

import (
    "github.com/fsnotify/fsnotify"
    "qqbot-reconstruction/internal/pkg/util"
    "qqbot-reconstruction/internal/pkg/variable"
)

func RegistryPlugins() *PluginEngine {
	// 加载注册表
	path := variable.GetConfigWd() + "plugins.yml"
	result := variable.ReadConfigs(path, &variable.PluginsConfig{})
	// 加载插件
	pluginEngine := initPluginEngine(result)
	// 热加载
	go util.WatchFile(path, func(e fsnotify.Event) {
		cfg := variable.ReadConfigs(path, &variable.PluginsConfig{})
		pluginEngine.Init(cfg, nil)
	})
	return pluginEngine
}
func initPluginRegistry(plugins []variable.PluginInfo) *PluginRegistry {
	pluginRegistry := NewPluginRegistry()
	// 插件自动扫描
	pluginRegistry.PluginScanner(plugins)
	return pluginRegistry
}

func initPluginEngine(plugins *variable.PluginsConfig) *PluginEngine {
	pluginEngine := NewPluginEngine()
	pluginEngine.Init(plugins, initPluginRegistry(plugins.Plugins))

	return pluginEngine
}
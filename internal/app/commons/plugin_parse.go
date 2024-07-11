
package commons

import (
    "qqbot-reconstruction/internal/app/plugins"
    "qqbot-reconstruction/internal/pkg/variable"
    "reflect"
)

func (r *PluginRegistry)PluginSanner(plugin *variable.PluginInfo)  {
    r.Register(plugin.Name, reflect.TypeOf(plugins.TestPlugin{}))
    r.Register(plugin.Name, reflect.TypeOf(plugins.MusicPlugin{}))
    r.Register(plugin.Name, reflect.TypeOf(plugins.MagnetPlugin{}))
    r.Register(plugin.Name, reflect.TypeOf(plugins.AliSearchPlugin{}))
}
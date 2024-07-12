
package commons

import (
    "qqbot-reconstruction/internal/app/plugins"
    "qqbot-reconstruction/internal/pkg/variable"
    "reflect"
)

func (r *PluginRegistry)PluginScanner(plugin []variable.PluginInfo)  {
    r.Register(plugin[0].Name, reflect.TypeOf(plugins.TestPlugin{}))
    r.Register(plugin[1].Name, reflect.TypeOf(plugins.MusicPlugin{}))
    r.Register(plugin[2].Name, reflect.TypeOf(plugins.MagnetPlugin{}))
    r.Register(plugin[3].Name, reflect.TypeOf(plugins.AliSearchPlugin{}))
}
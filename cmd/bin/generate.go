package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "os"
    "strings"
    "unicode"
)

var (
    head = `
package commons

import (
    "qqbot-reconstruction/internal/app/plugins"
    "qqbot-reconstruction/internal/pkg/variable"
    "reflect"
)

func (r *PluginRegistry)PluginScanner(plugin []variable.PluginInfo)  {
`
    end = `}`
    dir string
)

type PluginInfo struct {
    Name      string
    Keyword   string
    Whitelist []string
}

type PluginsConfig struct {
    Plugins []PluginInfo
}

func init()  {
    dir, _ = os.Getwd()
}
func main() {
    
    file, err := os.ReadFile(dir+"/../../configs/plugins.yml")
    if err != nil {
        fmt.Println(fmt.Sprintf("plugins配置文件读取失败: "), err)
        os.Exit(1)
    }
    plugin := &PluginsConfig{}
    err = yaml.Unmarshal(file, plugin)
    if err != nil {
        fmt.Println("配置文件解析到struct失败", err)
        os.Exit(1)
    }
    pluginInit(plugin.Plugins)
}

func pluginInit(plugins []PluginInfo) {
    path := dir+"/../../internal/app/commons/plugin_parse.go"
    file, err := os.Create(path)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    // 关闭文件
    defer file.Close()
    var builder strings.Builder
    builder.WriteString(head)
    for key, plugin := range plugins {
        camelCase := camelCase(plugin.Name) + "Plugin"
        builder.WriteString(fmt.Sprintf("    r.Register(plugin[%d].Name, reflect.TypeOf(plugins.%s{}))\n", key,camelCase))
    }
    builder.WriteString(end)
    result := builder.String()
    // 写入内容
    _, err = file.WriteString(result)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
}
func camelCase(input string) string {
    var result string
    capitalizeNext := true

    for _, char := range input {
        if char == '_' {
            // 遇到下划线，设置下一个字母大写
            capitalizeNext = true
        } else {
            if capitalizeNext {
                result += string(unicode.ToUpper(char))
                capitalizeNext = false
            } else {
                result += string(unicode.ToLower(char))
            }
        }
    }

    return result
}

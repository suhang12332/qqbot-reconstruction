package variable

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"qqbot-reconstruction/internal/pkg/log"
	"reflect"
	"sync"
)

var (
	Help    []string
	Urls    *ApiUrl
	Actions *Action
	once    sync.Once
)

// ReadConfigs
// @description: 读取配置文件
func ReadConfigs[T any](path string, t *T) *T {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Error(fmt.Sprintf("%s配置文件读取失败: ", reflect.TypeOf(t).String()), err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(file, t)
	if err != nil {
		log.Error("配置文件解析到struct失败", err)
		os.Exit(1)
	}
	return t
}

// Load
// @description: 加载配置文件
func init() {
	once.Do(func() {
		Actions = ReadConfigs(GetConfigWd()+"api.yaml", &Action{})
		Urls = ReadConfigs(GetConfigWd()+"url.yaml", &ApiUrl{})
		Help = []string{"help", "帮助", "例子", "事例", "用法", "说明", "手册"}
	})
}

// GetConfigWd
// @description: 用于获取当前目录
// @return: 返回当前目录
func GetConfigWd() string {
	dir, _ := os.Getwd()
	return dir + "/configs/"
}

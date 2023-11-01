package variable

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"qqbot-reconstruction/internal/pkg/log"
	"reflect"
	"runtime"
	"sync"
)

var (
	Urls    *ApiUrl
	Actions *Action

	once    sync.Once
)

// readConfigs
// @description: 读取配置文件
func readConfigs[T any](path string, t *T) *T {
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
		configPath := getWd()
		Actions = readConfigs(configPath+"/../../../configs/api.yaml", &Action{})
		Urls = readConfigs(configPath+"/../../../configs/url.yaml", &ApiUrl{})
	})

}

// getWd 
// @description: 用于获取当前目录
// @return: 返回当前目录
func getWd() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

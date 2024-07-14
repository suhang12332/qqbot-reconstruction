package variable

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"qqbot-reconstruction/internal/pkg/log"
	"qqbot-reconstruction/internal/pkg/util"
	"reflect"
	"sync"
)

var (
	Help    []string
	Urls    *ApiUrl
	Actions *Action

	Tips *Tip
	once sync.Once
)

// ReadConfigs
// @description: 读取配置文件
func ReadConfigs[T any](path string, t *T) *T {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf(fmt.Sprintf("%s配置文件读取失败: %v", reflect.TypeOf(t).String(), err))
	}
	if err = yaml.Unmarshal(file, t); err != nil {
		log.Fatal("配置文件解析到struct失败: %v", err)
	}
	return t
}

// Load
// @description: 加载配置文件
func init() {
	once.Do(func() {
		Actions = ReadConfigs(GetConfigWd()+"api.yml", &Action{})
		Urls = ReadConfigs(GetConfigWd()+"url.yml", &ApiUrl{})
		Tips = ReadConfigs(GetConfigWd()+"tips.yml", &Tip{})
		if Tips.MessageType == "img" {
			Tips.Info.NoArgs = util.TextParseImg(Tips.Info.NoArgs)
			Tips.Info.NoPermissions = util.TextParseImg(Tips.Info.NoPermissions)
		}
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

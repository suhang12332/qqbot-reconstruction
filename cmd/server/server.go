package main

import (
	"qqbot-reconstruction/internal/app/commons"
	client "qqbot-reconstruction/internal/app/server"
	"qqbot-reconstruction/internal/pkg/server"
	"qqbot-reconstruction/internal/pkg/util"
)

func main() {
	// 启动自定义的服务
	util.SafeGo(func() {
		server.StartHappyServer()
	})
	//注册插件
	plugins := commons.RegistryPlugins()
//	 client.Start(*plugins)

	client.Login(*plugins)

}

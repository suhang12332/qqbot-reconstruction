package main

import (
	"qqbot-reconstruction/internal/app/server"
	"qqbot-reconstruction/internal/pkg/server"
)

func main() {
	go func() {
		server.StartHappyServer()
	}()
	client.Start()
}

package main

import (
    "qqbot-reconstruction/internal/app/server"
)

func main() {
    go func() {
        server.StartHappyServer()
    }()
    server.Start()
}

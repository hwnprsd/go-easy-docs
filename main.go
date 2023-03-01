package main

import (
	"github.com/hwnprsd/go-api-docs/server"
)

func main() {
	forever := make(chan bool)
	go server.Run()
	<-forever
}

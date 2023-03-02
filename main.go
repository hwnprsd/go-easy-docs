package main

import (
	"github.com/hwnprsd/go-easy-docs/server"
)

func main() {
	forever := make(chan bool)
	go server.Run()
	<-forever
}

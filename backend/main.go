package main

import (
	"finance-go/adaptor/inbound"
	_ "finance-go/adaptor/outbound"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		inbound.Run()
	}()

	waitExist()
}

func waitExist() {
	signChannel := make(chan os.Signal, 1)
	signal.Notify(signChannel, syscall.SIGINT, syscall.SIGTERM)

	<-signChannel
	fmt.Print("程序退出...\n")
}

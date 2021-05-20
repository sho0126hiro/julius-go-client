package main

import (
	"github.com/sho0126hiro/julius-go-client/client"
	_ "github.com/sho0126hiro/julius-go-client/example/handler"
)

func main() {
	network := "tcp"
	ip := "localhost"
	port := "10500"
	c, err := client.NewClient(network, ip+":"+port)
	if err != nil {
		panic(err)
	}
	c.Run()
}

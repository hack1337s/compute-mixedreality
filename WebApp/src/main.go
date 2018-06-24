package main 

import (
	"fmt"
	"os"
	"serv"
	"client"
)

func main() {
	command := os.Args[1]
	switch command {
	case "bootserv":
		serv.Bootserv()
	case "bootclient":
		client.Bootclient()
	default:
		fmt.Printf("Not a valid command: use 'bootserv' or 'bootclient' instead")
	}
}
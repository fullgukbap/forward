package main

import (
	"log"
	"net"
)

func main() {
	in, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("occured error: ", err)
	}

}

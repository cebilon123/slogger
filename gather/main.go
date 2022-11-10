package main

import (
	"github.com/cebilon123/slogger/gather/server"
	"log"
)

func main() {
	handler := server.NewHandler()
	if err := handler.Start(); err != nil {
		log.Fatalln(err)
	}
}

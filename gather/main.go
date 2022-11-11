package main

import (
	"github.com/cebilon123/slogger/gather/server"
	"log"
	"runtime"
)

func main() {
	log.Println("Runtime cpus: ", runtime.NumCPU())
	handler := server.NewHandler()
	if err := handler.Start(); err != nil {
		log.Fatalln(err)
	}
}

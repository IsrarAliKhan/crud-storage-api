package main

import (
	"crud-storage-api/config"
	"crud-storage-api/internal/server"
	"log"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// server
	port, _ := strconv.Atoi(config.HttpPort)
	err := server.New(port).Start()
	if err != nil {
		log.Fatalf("start server err: %v", err)
	}
}

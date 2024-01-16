package main

import (
	"log"
	"webpDAV/internal/file"
	"webpDAV/internal/server"
)

func main() {

	server := server.NewServer()

	// starting file watcher used for converting images
	go file.SetupFileWatcher()

	log.Println("Starting webpDAV server...")
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}

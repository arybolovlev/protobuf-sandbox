package main

import (
	"flag"
	"log"

	"github.com/arybolovlev/protobuf-sandbox/client"
	"github.com/arybolovlev/protobuf-sandbox/server"
)

func main() {
	// Server options
	var srv bool
	flag.BoolVar(&srv, "server", false, "Run in the server mode")

	// Client options
	var cli bool
	flag.BoolVar(&cli, "client", false, "Run in the client mode")

	flag.Parse()

	if srv {
		log.Println("Running Server")
		server.Run()
	}

	if cli {
		log.Println("Running Client")
		client.Run()
	}
}

package main

import (
	"github.com/charmbracelet/log"
	"github.com/genof420/ferremas-api/internal/config"
	"github.com/genof420/ferremas-api/internal/server"
)

func main() {
	srv, err := server.New(config.Get())
	if err != nil {
		log.Fatal("Failed creating Server", "err", err)
		return
	}

	if err = srv.Start(); err != nil {
		log.Fatal("Failed running Server", "err", err)
		return
	}
}

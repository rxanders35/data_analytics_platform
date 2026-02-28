package main

import (
	"log"
	"os"

	"github.com/reidx/dap/controlplane/api-server/internal/router"
)

func main() {
	addr := os.Getenv("API_SERVER_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	router := router.NewRouter()

	log.Printf("api-server listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}

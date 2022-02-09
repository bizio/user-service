package main

import (
	"log"

	"github.com/bizio/user-service/pkg/cmd/server"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		log.Fatalf(" [Main] %v\n", err)
	}
}

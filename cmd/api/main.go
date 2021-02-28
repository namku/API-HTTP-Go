package main

import (
	"log"

	"github.com/namku/API-HTTP-Go/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}

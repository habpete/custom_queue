package main

import (
	"log"

	"github.com/custom_queue/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

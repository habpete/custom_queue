package main

import (
	"log"

	"github.com/custom_queue/internal/app"
	"github.com/custom_queue/internal/storage"
)

func main() {
	strg, err := storage.New(&storage.ConnectParams{})
	if err != nil {
		log.Fatal(err)
	}

	srv := app.NewService(strg)

	if err := app.Start(srv); err != nil {
		log.Fatal(err)
	}
}

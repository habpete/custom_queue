package main

import (
	"log"

	"golang.org/x/net/context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3000)
	defer cancel()

	if err := migrate(ctx, ""); err != nil {
		log.Fatal(err)
	}
}

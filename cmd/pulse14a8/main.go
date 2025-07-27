// cmd/pulse14a8/main.go
package main

import (
	"flag"
	"log"
	"os"

	"pulse14a8/internal/pulse14a8"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := pulse14a8.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

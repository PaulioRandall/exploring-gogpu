package main

import (
	"log"

	"github.com/PaulioRandall/exploring-gogpu/example1/app"
)

func main() {
	ea := app.NewExampleApp()

	if err := ea.Run(); err != nil {
		log.Fatal(err)
	}
}

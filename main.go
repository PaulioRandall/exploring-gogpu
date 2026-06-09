package main

import (
	"log"

	"github.com/PaulioRandall/exploring-gogpu/example_app"
)

func main() {
	ea := example_app.NewExampleApp()

	if err := ea.Run(); err != nil {
		log.Fatal(err)
	}
}

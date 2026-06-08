package main

import (
	"log"

	"github.com/PaulioRandall/exploring-gogpu/app"
)

func main() {
	a := app.NewApplication()
	defer a.Close()

	if e := a.Run(); e != nil {
		log.Fatal(e)
	}
}

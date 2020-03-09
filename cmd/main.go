package main

import (
	"log"

	app "github.com/comfortablynumb/goginrestapi/internal/app"
)

func main() {
	application, err := app.NewAppFromEnv()

	if err != nil {
		log.Fatal(err)
	}

	err = application.Run()

	if err != nil {
		log.Fatal(err)
	}
}

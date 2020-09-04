package main

import (

    "log"

	"github.com/rh-eu/golang-example-for-testing-che/pkg/app"
	"github.com/rh-eu/golang-example-for-testing-che/pkg/version"
)


func main() {
	app := app.NewApp()

	log.Printf("Starting goservd version: %v", version.VERSION)

	app.Run()
}
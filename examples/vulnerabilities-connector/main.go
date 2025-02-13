package main

import (
	"flag"
	"log"

	"vulnerabilities-connector/app"
)

func main() {
	var provider app.VulnerabilitiesConnectorProvider
	strProvider := flag.String("provider", "", "Provider to use for the example")
	flag.Parse()

	switch *strProvider {
	case string(app.CROUD_STRIKE):
		provider = app.CROUD_STRIKE
	case string(app.NUCLEUS):
		provider = app.NUCLEUS
	case string(app.QUALYS):
		provider = app.QUALYS
	case string(app.RAPID7):
		provider = app.RAPID7
	case string(app.TANIUM):
		provider = app.TANIUM
	case string(app.TENABLE):
		provider = app.TENABLE
	default:
		log.Fatalf("Please provide a valid provider: %s, %s, %s, %s or %s", app.CROUD_STRIKE, app.QUALYS, app.RAPID7, app.TANIUM, app.TENABLE)
	}

	app := app.VulnerabilitiesConnector{}
	app.Initialize(provider)

	app.Run()
}

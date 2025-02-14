package main

import (
	"flag"
	"log"

	appM "vulnerabilities-connector/app"
)

func main() {
	strProvider := flag.String("provider", "ALL", "Provider to use for the example")
	flag.Parse()

	var providers []appM.VulnerabilitiesConnectorProvider = []appM.VulnerabilitiesConnectorProvider{
		appM.CROUD_STRIKE,
		appM.NUCLEUS,
		appM.QUALYS,
		appM.RAPID7,
		appM.TANIUM,
		appM.TENABLE,
	}

	if *strProvider != "ALL" {
		providers = []appM.VulnerabilitiesConnectorProvider{}

		switch *strProvider {
		case string(appM.CROUD_STRIKE):
			providers = append(providers, appM.CROUD_STRIKE)
		case string(appM.NUCLEUS):
			providers = append(providers, appM.NUCLEUS)
		case string(appM.QUALYS):
			providers = append(providers, appM.QUALYS)
		case string(appM.RAPID7):
			providers = append(providers, appM.RAPID7)
		case string(appM.TANIUM):
			providers = append(providers, appM.TANIUM)
		case string(appM.TENABLE):
			providers = append(providers, appM.TENABLE)
		default:
			log.Fatalf("Please provide a valid provider: %s, %s, %s, %s, %s or %s", appM.CROUD_STRIKE, appM.NUCLEUS, appM.QUALYS, appM.RAPID7, appM.TANIUM, appM.TENABLE)
		}
	}

	app := appM.VulnerabilitiesConnector{}
	app.Initialize(providers)
	app.Run()
}

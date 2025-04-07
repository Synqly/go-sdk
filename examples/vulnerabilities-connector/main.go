package main

import (
	"flag"

	app "examples/vulnerabilities-connector/app"
)

func main() {
	boolCleanup := flag.Bool("cleanup", true, "Flag that execute a cleanup process before run")
	flag.Parse()

	app.NewVulnerabilitiesConnector().Start(*boolCleanup)
}

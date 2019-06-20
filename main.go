package main

import (
	"flag"

)

var credPath = flag.String("config", "config.json", "Path to configuration file containing the application's credentials.")
var lang = "en"

func main() {
	flag.Parse()

}


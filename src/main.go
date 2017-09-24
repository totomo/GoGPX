package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	destination, departure := getPlaces()

	key := getAPIKey()

	client := Client{
		APIKey:      key,
		Departure:   destination,
		Destination: departure,
	}

	response, error := client.Request()

	if error != nil {
		log.Fatalln(error)
	}

	printGPX(*response)

}

func getAPIKey() string {
	key := os.Getenv("APIKEY")

	if len(key) == 0 {
		log.Fatalln("APIKEY is empty.")
	}
	return key
}

func getPlaces() (string, string) {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
			Usage of %s:
				%s [OPTIONS] ARGS ...
			Options\n`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	startOption := flag.String("s", "", "Departure place name")
	goalOption := flag.String("g", "", "Destination place name")

	flag.Parse()

	if len(*startOption) == 0 || len(*goalOption) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	return *startOption, *goalOption

}

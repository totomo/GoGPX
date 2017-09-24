package main

import (
	"log"
	"os"
)

func main() {

	key := os.Getenv("APIKEY")

	if len(key) == 0 {
		log.Fatalln("APIKEY is empty.")
	}

	client := Client{
		APIKey:      key,
		Departure:   "名古屋駅",
		Destination: "名古屋城",
	}

	response, error := client.Request()

	if error != nil {
		log.Fatalln(error)
	}

	printGPX(*response)

}

package main

import "fmt"

func printGPX(response Entity) {

	if len(response.Routes) == 0 || len(response.Routes[0].Legs) == 0 {
		return
	}

	fmt.Println("<?xml version=\"1.0\"?>")
	fmt.Println("<gpx version=\"1.1\" creator=\"totomo\">")

	for _, value := range response.Routes[0].Legs[0].Steps {
		fmt.Printf("<wpt lat=%.8f lon=%.8f />\n", value.Start.Latitude, value.Start.Longitude)
	}

	fmt.Println("</gpx>")

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	APIKey      string
	Departure   string
	Destination string
}

func (c Client) Request() (*Entity, error) {

	parameters := url.Values{
		"origin":      {c.Departure},
		"destination": {c.Destination},
		"key":         {c.APIKey},
		"mode":        {"walking"},
	}

	url := "https://maps.googleapis.com/maps/api/directions/json"

	response, error := http.Get(url + "?" + parameters.Encode())
	if error != nil {
		return nil, error
	}
	fmt.Println(url + "?" + parameters.Encode())
	decoder := json.NewDecoder(response.Body)
	var data Entity
	error = decoder.Decode(&data)
	if error != nil {
		return nil, error
	}
	return &data, nil
}

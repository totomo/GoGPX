package main

type Entity struct {
	Routes []Routes `json:"routes"`
}

type Routes struct {
	Legs []Legs `json:"legs"`
}
type Legs struct {
	Steps []Steps `json:"steps"`
}
type Steps struct {
	Start Coordinate `json:"start_location"`
	End   Coordinate `json:"end_location"`
}

type Coordinate struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

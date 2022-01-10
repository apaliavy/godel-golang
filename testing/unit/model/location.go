package model

type LatLng struct {
	Latitude  float64
	Longitude float64
}

type Location struct {
	DisplayAddress string
	LatLng         *LatLng
}

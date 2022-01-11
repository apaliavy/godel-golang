package client

import (
	"errors"
	"math"

	"github.com/apaliavy/godel-golang/testing/unit/model"
)

var (
	ErrOriginNotProvided       = errors.New("nil origin is not allowed")
	ErrDestinationsNotProvided = errors.New("nil destination is not allowed")
)

// GetHaversineDistance returns distance in meters between two coords
func GetHaversineDistance(from, to *model.LatLng) (float64, error) {
	if from == nil {
		return 0, ErrOriginNotProvided
	}

	if to == nil {
		return 0, ErrDestinationsNotProvided
	}

	hsin := func(theta float64) float64 {
		return math.Pow(math.Sin(theta/2), 2)
	}

	var lat1r, lon1r, lat2r, lon2r float64
	lat1r = from.Latitude * math.Pi / 180
	lon1r = from.Longitude * math.Pi / 180
	lat2r = to.Latitude * math.Pi / 180
	lon2r = to.Longitude * math.Pi / 180

	const r = 6378100

	h := hsin(lat2r-lat1r) + math.Cos(lat1r)*math.Cos(lat2r)*hsin(lon2r-lon1r)

	return 2 * r * math.Asin(math.Sqrt(h)), nil
}

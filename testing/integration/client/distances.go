package client

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/sirupsen/logrus"

	"github.com/apaliavy/godel-golang/testing/integration/model"
	"github.com/apaliavy/godel-golang/testing/integration/repository"
)

var (
	ErrNotFoundInCache = errors.New("not found in cache")
)

type DistancesHandler struct {
	distancesCache repository.DistancesCache
}

func NewDistancesClient(dc repository.DistancesCache) *DistancesHandler {
	return &DistancesHandler{
		distancesCache: dc,
	}
}

func (h *DistancesHandler) GetDistance(ctx context.Context, from, to *model.Location) (float64, error) {
	var logger = logrus.New()

	if from == nil || from.LatLng == nil {
		return 0, fmt.Errorf("origin lat/lng is not provided")
	}

	if to == nil || to.LatLng == nil {
		return 0, fmt.Errorf("destination lat/lng is not provided")
	}

	distance, err := h.distancesCache.Get(ctx, from, to)
	if err == nil {
		return distance, nil
	} else if !errors.Is(err, ErrNotFoundInCache) {
		logger.WithError(err).Error("failed to fetch distance from cache")
	}

	distance, err = getHaversineDistance(from.LatLng, to.LatLng)
	if err != nil {
		return 0, err
	}

	if err := h.distancesCache.Put(ctx, from, to, distance); err != nil {
		logger.WithError(err).Error("failed to put data into cache")
	}

	return distance, nil
}

// getHaversineDistance returns distance in meters between two coords
func getHaversineDistance(from, to *model.LatLng) (float64, error) {
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

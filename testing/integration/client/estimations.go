package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/apaliavy/godel-golang/testing/integration/model"
)

type DistanceClient interface {
	GetDistance(ctx context.Context, from, to *model.Location) (float64, error)
}

type PricingStorage interface {
	GetPricingOptions() (*model.PricingOptions, error)
}

func CalculateEstimatedPrice(dc DistanceClient, s PricingStorage, from, to *model.Location) (float64, error) {
	distance, err := dc.GetDistance(context.Background(), from, to)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get distance from distances calculator")
	}

	pricingOptions, err := s.GetPricingOptions()
	if err != nil {
		return 0, errors.Wrap(err, "failed to retrieve pricing options")
	}

	return distance * pricingOptions.CostPerKm / 1000, nil
}

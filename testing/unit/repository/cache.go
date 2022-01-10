package repository

import (
	"context"

	"github.com/apaliavy/godel-golang/testing/unit/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 --fake-name DistancesCacheMock -o ../tools/testing/mocks/repository/cache.go . DistancesCache
type DistancesCache interface {
	Get(ctx context.Context, from, to *model.Location) (distance float64, err error)
	Put(ctx context.Context, from, to *model.Location, distance float64) error
}

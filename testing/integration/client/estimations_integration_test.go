// +build integration

package client_test

import (
	"context"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/apaliavy/godel-golang/testing/integration/storage"

	"github.com/apaliavy/godel-golang/testing/integration/client"

	"github.com/apaliavy/godel-golang/testing/integration/model"

	"testing"
)

type StubDistancesCalc struct {
}

func (dc *StubDistancesCalc) GetDistance(ctx context.Context, from, to *model.Location) (float64, error) {
	return 10.5, nil
}

func TestCalculateEstimatedPrice(t *testing.T) {
	distancesCalc := &StubDistancesCalc{}

	// bad practice, don't do like that in the real project
	p, err := storage.NewPricingOptions("postgresql://postgres:mysecretpassword@localhost:5432/prices?sslmode=disable")
	require.NoError(t, err)

	price, err := client.CalculateEstimatedPrice(
		distancesCalc,
		p,
		&model.Location{LatLng: &model.LatLng{22.3, 11.2}},
		&model.Location{LatLng: &model.LatLng{22.5, 11.9}},
	)

	require.NoError(t, err)

	assert.Equal(t, 0.21525, price)
}

package client_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apaliavy/godel-golang/testing/unit/client"
	"github.com/apaliavy/godel-golang/testing/unit/model"
	"github.com/apaliavy/godel-golang/testing/unit/repository"
	mockRepo "github.com/apaliavy/godel-golang/testing/unit/tools/testing/mocks/repository"
)

func TestDistancesHandler_GetDistance(t *testing.T) {
	cases := []struct {
		name        string
		from, to    *model.Location
		expectError bool
		err         error
		expected    float64
		getCache    func() repository.DistancesCache
	}{
		{
			name:        "validation error - origin is not provided",
			from:        nil,
			expectError: true,
			err:         fmt.Errorf("origin lat/lng is not provided"),
			getCache: func() repository.DistancesCache {
				dc := &mockRepo.DistancesCacheMock{}
				return dc
			},
		},
		{
			name:        "validation error - destination is not provided",
			from:        &model.Location{LatLng: &model.LatLng{}},
			expectError: true,
			err:         fmt.Errorf("destination lat/lng is not provided"),
			getCache: func() repository.DistancesCache {
				dc := &mockRepo.DistancesCacheMock{}
				return dc
			},
		},
		{
			name: "get distance between points - found in cache",
			from: &model.Location{LatLng: &model.LatLng{}},
			to:   &model.Location{LatLng: &model.LatLng{}},
			getCache: func() repository.DistancesCache {
				dc := &mockRepo.DistancesCacheMock{}
				dc.GetReturns(1001.25, nil)
				return dc
			},
			expectError: false,
			expected:    1001.25,
		},
		{
			name: "get distance between points - not found in cache",
			from: &model.Location{LatLng: &model.LatLng{
				Latitude:  20.10,
				Longitude: 20.20,
			}},
			to: &model.Location{LatLng: &model.LatLng{
				Latitude:  20.10,
				Longitude: 19.20,
			}},
			getCache: func() repository.DistancesCache {
				dc := &mockRepo.DistancesCacheMock{}
				dc.GetReturns(0, client.ErrNotFoundInCache)
				return dc
			},
			expectError: false,
			expected:    104538.73080187329,
		},
		{
			name: "get distance between points - an error received from cache",
			from: &model.Location{LatLng: &model.LatLng{
				Latitude:  20.10,
				Longitude: 20.20,
			}},
			to: &model.Location{LatLng: &model.LatLng{
				Latitude:  20.10,
				Longitude: 19.20,
			}},
			getCache: func() repository.DistancesCache {
				dc := &mockRepo.DistancesCacheMock{}
				dc.GetReturns(0, fmt.Errorf("an internal error"))
				return dc
			},
			expectError: false,
			expected:    104538.73080187329,
		},
		{
			name: "get distance between points - an error putting into cache",
			from: &model.Location{LatLng: &model.LatLng{}},
			to:   &model.Location{LatLng: &model.LatLng{}},
			getCache: func() repository.DistancesCache {
				dc := &mockRepo.DistancesCacheMock{}
				dc.GetReturns(1001.25, nil)
				dc.PutReturns(fmt.Errorf("failed to put data into cache"))
				return dc
			},
			expectError: false,
			expected:    1001.25,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			h := client.NewDistancesClient(tc.getCache())

			distance, err := h.GetDistance(context.Background(), tc.from, tc.to)

			if tc.expectError {
				require.Error(t, err)
				assert.EqualError(t, err, tc.err.Error())
				return
			}

			assert.Equal(t, tc.expected, distance)
		})
	}
}

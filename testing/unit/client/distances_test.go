package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apaliavy/godel-golang/testing/unit/client"
	"github.com/apaliavy/godel-golang/testing/unit/model"
)

func TestGetHaversineDistance(t *testing.T) {
	cases := []struct {
		name        string
		from, to    *model.LatLng
		expectError bool
		err         error
		expected    float64
	}{
		{
			name:        "origin is not provided - expect an error",
			from:        nil,
			to:          &model.LatLng{10.10, 20.20},
			expectError: true,
			err:         client.ErrOriginNotProvided,
		},
		{
			name:        "destination is not provided - expect an error",
			from:        &model.LatLng{10.10, 20.20},
			to:          nil,
			expectError: true,
			err:         client.ErrDestinationsNotProvided,
		},
		{
			name:        "same location given - zero distance",
			from:        &model.LatLng{10.10, 20.20},
			to:          &model.LatLng{10.10, 20.20},
			expectError: false,
			expected:    0,
		},
		{
			name:        "expect correct distance calculation",
			from:        &model.LatLng{20.10, 20.20},
			to:          &model.LatLng{20.10, 19.20},
			expectError: false,
			expected:    104538.73080187329,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			distance, err := client.GetHaversineDistance(tc.from, tc.to)

			if tc.expectError {
				require.Error(t, err)
				assert.EqualError(t, err, tc.err.Error())
				return
			}

			assert.Equal(t, tc.expected, distance)
		})
	}
}

package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type PricingOptions struct {
	CostPerKm float64 `json:"cost_per_km"`
}

func (o *PricingOptions) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *PricingOptions) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}

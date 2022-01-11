package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/apaliavy/godel-golang/testing/integration/model"
)

type Pricing struct {
	db *pgxpool.Pool
}

func NewPricingOptions(dns string) (*Pricing, error) {
	pool, err := pgxpool.Connect(context.Background(), dns)
	if err != nil {
		return nil, err
	}

	return &Pricing{db: pool}, nil
}

func (p *Pricing) GetPricingOptions() (options *model.PricingOptions, err error) {
	row := p.db.QueryRow(context.Background(), "SELECT configuration FROM pricing_options")
	return options, row.Scan(&options)
}

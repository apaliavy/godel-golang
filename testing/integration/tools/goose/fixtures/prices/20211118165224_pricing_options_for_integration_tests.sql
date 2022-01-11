-- +goose Up
-- +goose StatementBegin
INSERT INTO pricing_options(configuration) VALUES ('{"cost_per_km": 20.5}');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE pricing_options;
-- +goose StatementEnd

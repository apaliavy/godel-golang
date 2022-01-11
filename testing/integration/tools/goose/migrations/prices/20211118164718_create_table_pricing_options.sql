-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pricing_options (
    configuration jsonb
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pricing_options;
-- +goose StatementEnd


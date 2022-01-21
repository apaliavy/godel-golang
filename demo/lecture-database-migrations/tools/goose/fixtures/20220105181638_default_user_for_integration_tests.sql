-- +goose Up
-- +goose StatementBegin
INSERT INTO users(email, password, role) VALUES('alex@testuser.com', '$2a$10$Sfcitz5JncXQbZSWbxo/S.2.dNRA05aD7Ef.QFdqunbYzGSfHtwDu', 'ROLE_USER');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE users;
-- +goose StatementEnd

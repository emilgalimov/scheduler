-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks
(
    id          BIGSERIAL PRIMARY KEY,
    name        varchar,
    description text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd

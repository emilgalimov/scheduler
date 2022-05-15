-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_stages
(
    id                 BIGSERIAL PRIMARY KEY,
    name               varchar,
    description        text,
    minutes_from_start bigint,
    duration_minutes   bigint,
    task_id            bigint REFERENCES tasks (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_stages;
-- +goose StatementEnd

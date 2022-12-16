-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos
(
    id    int NOT NULL,
    title varchar,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd

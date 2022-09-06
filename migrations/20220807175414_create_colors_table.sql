-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.colors
(
    id   bigserial PRIMARY KEY,
    name varchar(100) UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.colors;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.apples
(
    id       bigserial PRIMARY KEY,
    color_id bigint NOT NULL,
    price    float8,
    CONSTRAINT fk_color
        FOREIGN KEY(color_id)
            REFERENCES colors(id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.apples;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users"
(
    "id"         BIGSERIAL                NOT NULL,

    "username"   VARCHAR                  NOT NULL,
    "password"   VARCHAR                  NOT NULL,
    "hashed"     BOOLEAN                  NOT NULL,

    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE NULL
);

ALTER TABLE IF EXISTS "users"
    ADD CONSTRAINT "PK_T_USR_C_ID"
        PRIMARY KEY ("id");

ALTER TABLE IF EXISTS "users"
    ADD CONSTRAINT "UK_T_USR_C_USR"
        UNIQUE ("username");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd
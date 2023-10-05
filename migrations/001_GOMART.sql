-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "sessions"
(
    "id"         VARCHAR                  NOT NULL,
    "user_id"    BIGINT                   NOT NULL,
    "token"      VARCHAR                  NOT NULL,
    "expires_at" TIMESTAMP WITH TIME ZONE NOT NULL
);

ALTER TABLE IF EXISTS "sessions"
    ADD CONSTRAINT "PK_T_SS_C_ID"
        PRIMARY KEY ("id");

ALTER TABLE IF EXISTS "sessions"
    ADD CONSTRAINT "UK_T_SS_C_ID_C_USR"
        UNIQUE ("id", "user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "sessions";
-- +goose StatementEnd
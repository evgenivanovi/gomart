-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "withdrawals"
(
    "id"         BIGSERIAL                NOT NULL,

    "user_id"    BIGINT                   NOT NULL,
    "order"      VARCHAR                  NOT NULL,
    "amount"     NUMERIC                  NOT NULL,

    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE NULL
);

ALTER TABLE IF EXISTS "withdrawals"
    ADD CONSTRAINT "PK_T_WDLS_C_ID"
        PRIMARY KEY ("id");

ALTER TABLE IF EXISTS "withdrawals"
    ADD CONSTRAINT "FK_T_WDLS_C_USR_ID"
        FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE IF EXISTS "withdrawals"
    ADD CONSTRAINT "UK_T_WDLS_C_ID_C_USR_ID"
        UNIQUE ("id", "user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "withdrawals";
-- +goose StatementEnd
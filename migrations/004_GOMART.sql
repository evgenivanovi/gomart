-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "balances"
(
    "id"      BIGSERIAL NOT NULL,
    "user_id" BIGINT    NOT NULL,
    "amount"  NUMERIC   NOT NULL
);

ALTER TABLE IF EXISTS "balances"
    ADD CONSTRAINT "PK_T_BLN_C_ID"
        PRIMARY KEY ("id");

ALTER TABLE IF EXISTS "balances"
    ADD CONSTRAINT "FK_T_BLN_C_USR_ID"
        FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE IF EXISTS "balances"
    ADD CONSTRAINT "UK_T_BLN_C_ID_C_USR_ID"
        UNIQUE ("id", "user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "balances";
-- +goose StatementEnd
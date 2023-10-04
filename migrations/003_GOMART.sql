-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "orders"
(
    "id"         VARCHAR                  NOT NULL,
    "user_id"    BIGINT                   NOT NULL,

    "status"     VARCHAR                  NOT NULL,
    "accrual"    NUMERIC                  NOT NULL,

    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE NULL,
    "deleted_at" TIMESTAMP WITH TIME ZONE NULL
);

ALTER TABLE IF EXISTS "orders"
    ADD CONSTRAINT "PK_T_ORD_C_ID"
        PRIMARY KEY ("id");

ALTER TABLE IF EXISTS "orders"
    ADD CONSTRAINT "FK_T_ORD_C_USR_ID"
        FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE IF EXISTS "orders"
    ADD CONSTRAINT "UK_T_ORD_C_ID_C_USR_ID"
        UNIQUE ("id", "user_id");

ALTER TABLE IF EXISTS "orders"
    ADD CONSTRAINT "CK_T_ORD_C_STS"
        CHECK ("status" in ('NEW', 'REGISTERED', 'PROCESSING', 'PROCESSED', 'INVALID'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "orders";
-- +goose StatementEnd
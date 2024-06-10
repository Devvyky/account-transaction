CREATE TABLE "accounts" (
  "account_id" bigserial PRIMARY KEY,
  "document_number" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "operation_types" (
  "operation_type_id" bigserial PRIMARY KEY,
  "description" varchar NOT NULL
);

CREATE TABLE "transactions" (
  "transaction_id" bigserial PRIMARY KEY,
  "operation_type_id" bigint NOT NULL,
  "account_id" bigint NOT NULL,
  "amount" numeric(15,2) NOT NULL,
  "event_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("account_id");

COMMENT ON COLUMN "transactions"."amount" IS 'can be negative or positive';

ALTER TABLE "transactions" ADD FOREIGN KEY ("operation_type_id") REFERENCES "operation_types" ("operation_type_id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

INSERT INTO operation_types (description) VALUES
('PURCHASE'),
('INSTALLMENT PURCHASE'),
('WITHDRAWAL'),
('PAYMENT');
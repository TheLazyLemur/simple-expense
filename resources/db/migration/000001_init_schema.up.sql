CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "expenses" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "invoices" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial NOT NULL,
  "expense_id" bigint NOT NULL,
  "url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("name", "email");

CREATE INDEX ON "expenses" ("owner");

CREATE INDEX ON "invoices" ("owner");

ALTER TABLE "expenses" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id");


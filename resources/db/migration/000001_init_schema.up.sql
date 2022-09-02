CREATE TABLE IF NOT EXISTS "organisations" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "owner" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "user_organisations_access" (
  "user_id" bigserial NOT NULL,
  "organisation_id" bigserial NOT NULL,
  UNIQUE ("user_id", "organisation_id")
);

CREATE TABLE IF NOT EXISTS "expenses" (
  "id" bigserial PRIMARY KEY,
  "uploader" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "organisation_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "invoices" (
  "id" bigserial PRIMARY KEY,
  "organisation_id" bigserial NOT NULL,
  "uploader" bigserial NOT NULL,
  "expense_id" bigint NOT NULL,
  "url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "organisations" ("owner");

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("name", "email");

CREATE INDEX ON "user_organisations_access" ("user_id");

CREATE INDEX ON "user_organisations_access" ("organisation_id");

CREATE INDEX ON "user_organisations_access" ("user_id", "organisation_id");

CREATE INDEX ON "expenses" ("uploader");

CREATE INDEX ON "expenses" ("organisation_id");

CREATE INDEX ON "expenses" ("organisation_id", "uploader");

CREATE INDEX ON "invoices" ("organisation_id");

CREATE INDEX ON "invoices" ("uploader");

CREATE INDEX ON "invoices" ("organisation_id", "uploader");

ALTER TABLE "organisations" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "user_organisations_access" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_organisations_access" ADD FOREIGN KEY ("organisation_id") REFERENCES "organisations" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("uploader") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("organisation_id") REFERENCES "organisations" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("organisation_id") REFERENCES "organisations" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("uploader") REFERENCES "users" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id");


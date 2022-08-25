ALTER TABLE IF EXISTS "users" ADD username varchar NOT NULL;

ALTER TABLE IF EXISTS "users" ADD password varchar NOT NULL;

ALTER TABLE IF EXISTS "users" ADD salt varchar NOT NULL;

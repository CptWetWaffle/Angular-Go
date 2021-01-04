CREATE TABLE "accounts" (
                           "id" SERIAL PRIMARY KEY,
                           "owner" varchar NOT NULL,
                           "balance" decimal NOT NULL,
                           "currency" varchar NOT NULL,
                           "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
                           "id" SERIAL PRIMARY KEY,
                           "account_id" bigserial NOT NULL,
                           "amount" decimal NOT NULL,
                           "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
                             "id" SERIAL PRIMARY KEY,
                             "from_account_id" bigserial NOT NULL,
                             "to_account_id" bigserial NOT NULL,
                             "amount" decimal NOT NULL,
                             "created_at" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

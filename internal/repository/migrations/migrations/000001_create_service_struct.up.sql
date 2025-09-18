CREATE TABLE "user" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar(50)
);

CREATE TABLE "service" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "amount" decimal(10,2) NOT NULL
);

CREATE TABLE "user_purchase" (
  "id" serial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "service_id" integer NOT NULL,
  "created_at" date NOT NULL DEFAULT (CURRENT_DATE),
  "end_date" date
);

ALTER TABLE "user_purchase" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_purchase" ADD FOREIGN KEY ("service_id") REFERENCES "service" ("id");

CREATE INDEX idx_user_purchase_user_id ON "user_purchase" ("user_id");

CREATE INDEX idx_service_name ON service ("name");

CREATE INDEX idx_user_purchase_created_at ON "user_purchase" ("created_at");
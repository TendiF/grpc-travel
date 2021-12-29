CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "country_code" varchar,
  "phone" varchar,
  "first_name" varchar,
  "last_name" varchar,
  "gender" varchar,
  "email" varchar,
  "birth_date" date,
  "password" varchar,
  "document_number" varchar,
  "address" varchar,
  "verification_code" varchar,
  "status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);
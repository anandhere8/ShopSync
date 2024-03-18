CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone_number" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
  CONSTRAINT valid_email CHECK (email ~* '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'),
  
  -- Check constraint to validate phone number format
  -- CONSTRAINT valid_phone_number CHECK (phone_number ~* '^\+?\d{1,3}[- ]?\d{3}[- ]?\d{3}[- ]?\d{4}$'),
  CONSTRAINT valid_indian_phone_number CHECK (phone_number ~ '^(?:\+?91|0)?[789]\d{9}$'),
  -- Check constraint to ensure password is at least 8 characters long
  CONSTRAINT password_length CHECK (LENGTH(password_hash) >= 8)
);

CREATE TABLE "shops" (
  "shop_id" bigserial PRIMARY KEY,
  "owner_id" bigint NOT NULL,
  "shop_name" varchar NOT NULL,
  "shop_description" text,
  "shop_address" text,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "employee" (
  "emp_id" bigserial PRIMARY KEY,
  "shop_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "roles" (
  "role_id" bigserial PRIMARY KEY,
  "role_name" varchar UNIQUE NOT NULL
);

CREATE TABLE "items" (
  "item_id" bigserial PRIMARY KEY,
  "owner_id" bigint NOT NULL,
  "shop_id" bigint NOT NULL,
  "item_code" varchar NOT NULL,
  "item_name" varchar NOT NULL,
  "description" varchar,
  "quantity" bigint NOT NULL DEFAULT 1,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_username" ON "users" ("username");

CREATE INDEX "idx_shop_name" ON "shops" ("shop_name");

CREATE INDEX "idx_owner_id" ON "shops" ("owner_id");

CREATE INDEX "idx_shop_id" ON "employee" ("shop_id");

CREATE INDEX "idx_role_id" ON "employee" ("role_id");

CREATE INDEX "idx_item_shop_id" ON "items" ("shop_id");

CREATE INDEX "idx_created_at" ON "items" ("created_at");

ALTER TABLE "shops" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("user_id");

ALTER TABLE "employee" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("shop_id");

ALTER TABLE "employee" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "employee" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "items" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("user_id");

ALTER TABLE "items" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("shop_id");
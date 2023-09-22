CREATE TABLE "branches" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar,
  "phone_number" varchar,
  "created_at" timestamp DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT (current_timestamp)
);

CREATE TABLE "categories" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "parent_id" uuid,
  "created_at" timestamp DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT (current_timestamp)
);

CREATE TABLE "products" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "price" numeric NOT NULL,
  "barcode" varchar UNIQUE NOT NULL,
  "category_id" uuid,
  "created_at" timestamp DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT (current_timestamp)
);

CREATE TABLE "coming_tables" (
  "id" uuid PRIMARY KEY,
  "coming_id" varchar NOT NULL,
  "branch_id" uuid,
  "date_time" timestamp,
  "status" varchar DEFAULT 'in_process',
  "created_at" timestamp DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT (current_timestamp)
);

CREATE TABLE "coming_table_products" (
  "id" uuid PRIMARY KEY,
  "category_id" uuid,
  "name" varchar NOT NULL,
  "price" numeric NOT NULL,
  "barcode" varchar UNIQUE NOT NULL,
  "count" numeric NOT NULL DEFAULT 0,
  "total_price" numeric DEFAULT 0,
  "coming_table_id" uuid,
  "created_at" timestamp DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT (current_timestamp)
);

CREATE TABLE "remainings" (
  "id" uuid PRIMARY KEY,
  "branch_id" uuid,
  "category_id" uuid,
  "name" varchar NOT NULL,
  "price" numeric NOT NULL,
  "barcode" varchar UNIQUE NOT NULL,
  "count" numeric NOT NULL DEFAULT 0,
  "total_price" numeric DEFAULT 0,
  "created_at" timestamp DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT (current_timestamp)
);

ALTER TABLE "categories" ADD FOREIGN KEY ("parent_id") REFERENCES "categories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "coming_tables" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "coming_table_products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "coming_table_products" ADD FOREIGN KEY ("coming_table_id") REFERENCES "coming_tables" ("id");

ALTER TABLE "remainings" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "remainings" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "full_name" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role" varchar NOT NULL DEFAULT 'user',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "histories" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "user_id" uuid,
  "al_id" varchar NOT NULL,
  "mangadex_id" varchar NOT NULL,
  "cover_image" varchar NOT NULL,
  "title" varchar NOT NULL,
  "reading_chapter" varchar NOT NULL,
  "path" varchar NOT NULL,
  "last_read_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "histories" ("mangadex_id", "user_agent");

CREATE UNIQUE INDEX ON "histories" ("mangadex_id", "user_id");

ALTER TABLE "histories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET NULL;

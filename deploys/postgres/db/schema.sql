CREATE TABLE "sessions" (
  "id" bigserial PRIMARY KEY,
  "session_id" text,
  "email" text,
  "refresh_token" text,
  "expires_at" text,
  "user_agent" text,
  "client_ip" text,
  "is_blocked" boolean DEFAULT false
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" text UNIQUE,
  "password" text,
  "active" boolean DEFAULT true
);

CREATE INDEX ON "sessions" ("session_id");

CREATE INDEX ON "sessions" ("email");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");

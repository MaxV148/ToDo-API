CREATE TABLE "todo" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "done" boolean NOT NULL DEFAULT false,
  "created_by" bigint NOT NULL,
  "category" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "user" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "todo" ("title");

CREATE INDEX ON "user" ("username");

CREATE INDEX ON "user" ("password");

CREATE INDEX ON "category" ("name");

CREATE INDEX ON "category" ("user");

ALTER TABLE "todo" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");

ALTER TABLE "todo" ADD FOREIGN KEY ("category") REFERENCES "category" ("id");

ALTER TABLE "category" ADD FOREIGN KEY ("user") REFERENCES "user" ("id");

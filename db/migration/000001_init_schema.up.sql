CREATE TABLE "Assets" (
  "id" bigserial PRIMARY KEY,
  "internal_id" bigserial UNIQUE NOT NULL,
  "asset_name" varchar UNIQUE NOT NULL,
  "asset_created_at" varchar NOT NULL,
  "status" bool NOT NULL DEFAULT false,
  "asset_link" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Assets" ("internal_id");

CREATE INDEX ON "Assets" ("asset_name");

CREATE INDEX ON "Assets" ("asset_created_at");

CREATE INDEX ON "Assets" ("status");
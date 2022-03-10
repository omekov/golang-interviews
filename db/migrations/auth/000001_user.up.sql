CREATE TABLE "permission" (
    "id" serial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "key" varchar(255) NOT NULL,
    "description" varchar(255) NOT NULL,
    "permission_id" bigint NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "deleted_at" timestamptz NULL
);

ALTER TABLE "permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");

CREATE INDEX ON "permission" ("permission_id");

COMMENT ON COLUMN "permission"."created_at" IS 'дата создание строки';
COMMENT ON COLUMN "permission"."updated_at" IS 'дата обновления строки';
COMMENT ON COLUMN "permission"."deleted_at" IS 'дата удаления строки';


CREATE TABLE "role" (
    "id" serial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "description" varchar(255) NOT NULL,
    "permission_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "deleted_at" timestamptz NULL
);

ALTER TABLE "role" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");

CREATE INDEX ON "role" ("permission_id");

COMMENT ON COLUMN "role"."created_at" IS 'дата создание строки';
COMMENT ON COLUMN "role"."updated_at" IS 'дата обновления строки';
COMMENT ON COLUMN "role"."deleted_at" IS 'дата удаления строки';

CREATE TABLE "user" (
    "id" serial PRIMARY KEY,
    "username" varchar(255) NOT NULL UNIQUE,
    "password"  varchar(255) NOT NULL UNIQUE,
    "firstname" varchar(255) NULL,
    "lastname" varchar(255) NULL,
    "phone_number" varchar(255) NULL,
    "role_id" bigint NOT NULL,
    "is_disabled" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "deleted_at" timestamptz NULL
);

ALTER TABLE "user" ADD FOREIGN KEY ("role_id") REFERENCES "role_id" ("id");

CREATE INDEX ON "user" ("role_id");

COMMENT ON COLUMN "user"."created_at" IS 'дата создание строки';
COMMENT ON COLUMN "user"."updated_at" IS 'дата обновления строки';
COMMENT ON COLUMN "user"."deleted_at" IS 'дата удаления строки';


-- create "users" table
CREATE TABLE "users" (
    "id" character varying NOT NULL, 
    "name" character varying NOT NULL, 
    "age" bigint NOT NULL, 
    "created_at" timestamptz NOT NULL, 
    "updated_at" timestamptz NOT NULL, 
    PRIMARY KEY ("id")
);

-- create "todos" table
CREATE TABLE "todos" (
    "id" character varying NOT NULL,
    "name" character varying NOT NULL DEFAULT '',
    "status" character varying NOT NULL DEFAULT 'IN_PROGRESS',
    "priority" bigint NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "user_id" character varying NULL,
    PRIMARY KEY ("id"), CONSTRAINT 
    "todos_users_todos" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET NULL
);

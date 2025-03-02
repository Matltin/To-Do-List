CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "create_at" timestamp DEFAULT(now()),
    "username" varchar NOT NULL,
    "hashed_password" varchar NOT NULL,
    "email" varchar NOT NULL
);


CREATE TABLE "todos" (
    "id" bigserial PRIMARY KEY,
    "user_id" int NOT NULL,
    "create_at" timestamp DEFAULT (now()),
    "update_at" timestamp NOT NULL,
    "finished_at" timestamp NOT NULL,
    "title" varchar NOT NULL,
    "description" varchar NOT NULL,
    "is_done" boolean DEFAULT false
);


ALTER TABLE "todos"
ADD CONSTRAINT fk_user FOREIGN KEY ("user_id") 
REFERENCES "users" ("id") ON DELETE RESTRICT;

CREATE INDEX todos_user_id_idx 
ON "todos" ("user_id");

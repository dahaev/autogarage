CREATE TABLE "cars" (
                        "id" bigserial PRIMARY KEY,
                        "model_name" varchar NOT NULL,
                        "equipment" varchar NOT NULL,
                        "color" varchar NOT NULL,
                        "country" varchar NOT NULL,
                        "price" bigint NOT NULL,
                        "valid" varchar NOT NULL,
                        "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "orders" (
                          "id" bigserial PRIMARY KEY,
                          "car" bigint NOT NULL,
                          "client" bigint NOT NULL,
                          "manager" bigint NOT NULL,
                          "magazine" bigint NOT NULL,
                          "delivery_time" int NOT NULL,
                          "car_price" bigint NOT NULL,
                          "delivery_price" bigint NOT NULL DEFAULT 0,
                          "tax" bigint NOT NULL DEFAULT 0,
                          "total_price" bigint NOT NULL
);

CREATE TABLE "clients" (
                           "id" bigserial PRIMARY KEY,
                           "name" varchar NOT NULL,
                           "country" varchar NOT NULL,
                           "phone_number" varchar NOT NULL
);

CREATE TABLE "managers" (
                            "id" bigserial PRIMARY KEY,
                            "name" varchar NOT NULL,
                            "town" varchar NOT NULL
);

CREATE TABLE "magazines" (
                             "id" bigserial PRIMARY KEY,
                             "address" varchar NOT NULL
);

CREATE INDEX ON "cars" ("model_name");

CREATE INDEX ON "cars" ("equipment");

ALTER TABLE "orders" ADD FOREIGN KEY ("car") REFERENCES "cars" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("client") REFERENCES "clients" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("manager") REFERENCES "managers" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("magazine") REFERENCES "magazines" ("id");

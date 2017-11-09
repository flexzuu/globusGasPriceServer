CREATE TABLE "gasPrices" (
    "id" serial,
    "lastUpdated" timestamptz NOT NULL,
    "e5" decimal NOT NULL,
    "e10" decimal NOT NULL,
    "superPlus" decimal NOT NULL,
    "diesel" decimal NOT NULL,
    "autogas" decimal NOT NULL,
    PRIMARY KEY ("id")
);

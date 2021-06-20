CREATE TABLE income (
    income_uuid UUID PRIMARY KEY,
    account_uuid UUID NOT NULL REFERENCES account ON DELETE RESTRICT,
    name TEXT NOT NULL,
    description TEXT,
    amount NUMERIC(12, 4) NOT NULL,
    date_made TIMESTAMPTZ NOT NULL
);
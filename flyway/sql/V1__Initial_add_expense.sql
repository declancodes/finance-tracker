CREATE TABLE expense_category (
    expense_category_uuid UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE expense (
    expense_uuid UUID PRIMARY KEY,
    expense_category_uuid UUID NOT NULL REFERENCES expense_category ON DELETE RESTRICT,
    name TEXT NOT NULL,
    description TEXT,
    amount NUMERIC(12, 4) NOT NULL,
    date_incurred TIMESTAMPTZ NOT NULL
);

CREATE TABLE expense_category (
    expense_category_uuid UUID PRIMARY KEY,
    name VARCHAR(25) UNIQUE NOT NULL,
    description VARCHAR(150)
);

CREATE TABLE expense (
    expense_uuid UUID PRIMARY KEY,
    expense_category_uuid UUID NOT NULL REFERENCES expense_category ON DELETE RESTRICT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(150),
    amount NUMERIC(12, 4) NOT NULL,
    date_incurred TIMESTAMPTZ NOT NULL
);

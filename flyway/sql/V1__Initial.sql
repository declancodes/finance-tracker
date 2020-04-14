CREATE TABLE expense_category (
    expense_category_uuid UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE expense (
    expense_uuid UUID PRIMARY KEY,
    expense_category_uuid UUID NOT NULL REFERENCES expense_category ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    amount NUMERIC(12, 4) NOT NULL,
    date_incurred TIMESTAMPTZ NOT NULL
);

CREATE TABLE account_category (
    account_category_uuid UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE account (
    account_uuid UUID PRIMARY KEY,
    account_category_uuid UUID NOT NULL REFERENCES account_category ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    amount NUMERIC(12, 4) NOT NULL
);

CREATE TABLE contribution (
    contribution_uuid UUID PRIMARY KEY,
    account_uuid UUID NOT NULL REFERENCES account ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    amount NUMERIC(12, 4) NOT NULL,
    date_made TIMESTAMPTZ NOT NULL
);

CREATE TABLE account_category (
    account_category_uuid UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE account (
    account_uuid UUID PRIMARY KEY,
    account_category_uuid UUID NOT NULL REFERENCES account_category ON DELETE RESTRICT,
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    amount NUMERIC(12, 4) NOT NULL
);

CREATE TABLE contribution (
    contribution_uuid UUID PRIMARY KEY,
    account_uuid UUID NOT NULL REFERENCES account ON DELETE RESTRICT,
    name TEXT NOT NULL,
    description TEXT,
    amount NUMERIC(12, 4) NOT NULL,
    date_made TIMESTAMPTZ NOT NULL
);

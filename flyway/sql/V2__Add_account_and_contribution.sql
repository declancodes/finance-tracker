CREATE TABLE account_category (
    account_category_uuid UUID PRIMARY KEY,
    name VARCHAR(25) UNIQUE NOT NULL,
    description VARCHAR(150)
);

CREATE TABLE account (
    account_uuid UUID PRIMARY KEY,
    account_category_uuid UUID NOT NULL REFERENCES account_category ON DELETE RESTRICT,
    name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(150),
    amount NUMERIC(12, 4) NOT NULL
);

CREATE TABLE contribution (
    contribution_uuid UUID PRIMARY KEY,
    account_uuid UUID NOT NULL REFERENCES account ON DELETE RESTRICT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(150),
    amount NUMERIC(12, 4) NOT NULL,
    date_made TIMESTAMPTZ NOT NULL
);

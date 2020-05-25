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

CREATE TABLE fund (
    fund_uuid UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    ticker_symbol VARCHAR(6) UNIQUE NOT NULL,
    share_price NUMERIC(12, 4) NOT NULL
);

CREATE TABLE holding (
    holding_uuid UUID PRIMARY KEY,
    account_uuid UUID NOT NULL REFERENCES account ON DELETE RESTRICT,
    fund_uuid UUID NOT NULL REFERENCES fund ON DELETE RESTRICT,
    shares NUMERIC(12, 4) NOT NULL,
    UNIQUE (account_uuid, fund_uuid)
);

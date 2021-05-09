CREATE TABLE asset_category (
    asset_category_uuid UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE fund (
    fund_uuid UUID PRIMARY KEY,
    asset_category_uuid UUID NOT NULL REFERENCES asset_category ON DELETE RESTRICT,
    name TEXT NOT NULL,
    ticker_symbol VARCHAR(5) NOT NULL,
    share_price NUMERIC(12, 4) NOT NULL,
    expense_ratio NUMERIC(12, 8) NOT NULL,
    is_private BOOLEAN NOT NULL
);

CREATE TABLE holding (
    holding_uuid UUID PRIMARY KEY,
    account_uuid UUID NOT NULL REFERENCES account ON DELETE RESTRICT,
    fund_uuid UUID NOT NULL REFERENCES fund ON DELETE RESTRICT,
    shares NUMERIC(12, 4) NOT NULL,
    UNIQUE (account_uuid, fund_uuid)
);

CREATE TABLE holding (
    holding_uuid UUID PRIMARY KEY,
    account_uuid UUID NOT NULL REFERENCES account ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    ticker_symbol VARCHAR(6),
    shares NUMERIC(12, 4) NOT NULL,
    UNIQUE (account_uuid, ticker_symbol)
);

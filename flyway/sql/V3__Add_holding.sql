CREATE TABLE holding (
    holding_uuid UUID PRIMARY KEY,
    account_uuid UUID NOT NULL REFERENCES account ON DELETE RESTRICT,
    name VARCHAR(255) UNIQUE NOT NULL,
    ticker_symbol VARCHAR(255),
    shares NUMERIC(12, 4) NOT NULL
);

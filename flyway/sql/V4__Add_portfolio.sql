CREATE TABLE portfolio (
    portfolio_uuid UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE portfolio_holding_mapping (
    portfolio_holding_mapping_uuid UUID PRIMARY KEY,
    portfolio_uuid UUID NOT NULL REFERENCES portfolio ON DELETE RESTRICT,
    holding_uuid UUID NOT NULL REFERENCES holding ON DELETE RESTRICT,
    UNIQUE (portfolio_uuid, holding_uuid)
);

CREATE TABLE portfolio_asset_category_mapping (
    portfolio_asset_category_mapping_uuid UUID PRIMARY KEY,
    portfolio_uuid UUID NOT NULL REFERENCES portfolio ON DELETE RESTRICT,
    asset_category_uuid UUID NOT NULL REFERENCES asset_category ON DELETE RESTRICT,
    percentage NUMERIC(12, 4) NOT NULL,
    UNIQUE (portfolio_uuid, asset_category_uuid)
);

CREATE TABLE IF NOT EXISTS accounts (
    account_number BIGSERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    balance DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- insert a sample record
INSERT INTO accounts (first_name, last_name, balance)
VALUES ('Alice', 'Doe', 1000.00);

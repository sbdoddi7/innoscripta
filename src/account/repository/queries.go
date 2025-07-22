package repository

const queryCreateAccount = `
    INSERT INTO accounts (first_name, last_name, balance)
    VALUES ($1, $2, $3)
    RETURNING account_number
`

const queryGetAccountByNumber = `
    SELECT account_number, first_name, last_name, balance, created_at
    FROM accounts
    WHERE account_number = $1
`

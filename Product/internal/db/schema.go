package db

var schema = `
CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    name text,
    description text,
    quantity int,
    price DECIMAL(6,2)
)
`

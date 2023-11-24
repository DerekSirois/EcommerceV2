package db

var schema = `
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username text,
    password text,
    email text,
    isAdmin boolean
)
`

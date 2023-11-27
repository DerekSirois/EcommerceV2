package db

var schema = `
CREATE TABLE IF NOT EXISTS cart(
    id SERIAL PRIMARY KEY,
    userId int
);

CREATE TABLE IF NOT EXISTS cartItem(
    id SERIAL PRIMARY KEY,
    cartId int,
    productId int,
    quantity int
)
`

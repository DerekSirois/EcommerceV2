package db

type Cart struct {
}

func AddItem(cartId int, productId int, quantity int) error {
	_, err := Db.Exec("INSERT INTO cartItem(cartId, productId, quantity) VALUES ($1, $2, $3)", cartId, productId, quantity)
	return err
}

func RemoveItem(cartId int, productId int) error {
	_, err := Db.Exec("DELETE FROM cartItem WHERE cartId = $1 and productId = $2", cartId, productId)
	return err
}

func UpdateItem(cartId int, productId int, quantity int) error {
	_, err := Db.Exec("UPDATE cartItem SET quantity = $1 WHERE cartId = $2 and productId = $3", quantity, cartId, productId)
	return err
}

func CreateCart(userId int) error {
	_, err := Db.Exec("INSERT INTO cart (userId) VALUES ($1)", userId)
	return err
}

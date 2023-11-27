package db

type Product struct {
	Id       int `db:"productId"`
	Quantity int
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

func ClearCart(cartId int) error {
	_, err := Db.Exec("DELETE FROM cartItem WHERE cartId = $1", cartId)
	return err
}

func GetAllProduct(cartId int) ([]*Product, error) {
	p := make([]*Product, 0)
	err := Db.Select(&p, "SELECT * FROM cartItem WHERE cartId = $1", cartId)
	return p, err
}

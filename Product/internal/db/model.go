package db

type Product struct {
	Id          int
	Name        string
	Description string
	Quantity    int
	Price       float64
}

func GetAll() ([]*Product, error) {
	p := make([]*Product, 0)
	err := Db.Select(p, "SELECT * FROM products")
	return p, err
}

func GetAllAvailable() ([]*Product, error) {
	p := make([]*Product, 0)
	err := Db.Select(p, "SELECT * FROM products WHERE quantity > 0")
	return p, err
}

func GetAllOutOfStock() ([]*Product, error) {
	p := make([]*Product, 0)
	err := Db.Select(p, "SELECT * FROM products WHERE quantity <= 0")
	return p, err
}

func GetById(id int) (*Product, error) {
	var p *Product
	err := Db.Get(p, "SELECT * FROM products WHERE id = $1", id)
	return p, err
}

func Create(p Product) error {
	_, err := Db.Exec("INSERT INTO products(name, description, quantity, price) VALUES ($1, $2, $3, $4)", p.Name, p.Description, p.Quantity, p.Price)
	return err
}

func Update(p Product, id int) error {
	_, err := Db.Exec("UPDATE products SET name = $1, description = $2, quantity = $3, price = $4 WHERE id = $5", p.Name, p.Description, p.Quantity, p.Price, id)
	return err
}

func Delete(id int) error {
	_, err := Db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}

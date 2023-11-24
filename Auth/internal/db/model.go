package db

type User struct {
	Id       int
	Username string
	Password string
	Email    string
	IsAdmin  bool
}

func GetByUsername(username string) (*User, error) {
	u := &User{}
	err := Db.Get(u, "SELECT * FROM users WHERE username = $1", username)
	return u, err
}

func Create(u User) error {
	_, err := Db.Exec("INSERT INTO users(username, password, email, isAdmin) VALUES ($1, $2, $3, $4)",
		u.Username, u.Password, u.Email, u.IsAdmin)
	return err
}

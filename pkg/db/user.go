package db

type UserDb struct{}

type User struct {
	Id   int64
	Name string
}

func (userDb *UserDb) GetUserById(id int64) (*User, error) {
	return &User{Id: id, Name: "dehwyy"}, nil
}

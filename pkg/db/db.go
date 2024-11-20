package db

type Database struct {
	userDb *UserDb
}

func NewDatabase() *Database {
	return &Database{userDb: &UserDb{}}
}

func (db *Database) User() *UserDb {
	return db.userDb
}

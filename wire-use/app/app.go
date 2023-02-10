package app

type User struct {
	ID   int
	Name string
}

type UserService struct {
	DB *Database
}

type Database struct {
	Users map[int]User
}

func NewUserService(db *Database) *UserService {
	return &UserService{DB: db}
}

func NewDatabase() *Database {
	return &Database{
		Users: map[int]User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
			3: {ID: 3, Name: "Charlie"},
		},
	}
}

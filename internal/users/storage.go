package users

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) geAllUsers() ([]user, error) {
	rows, err := u.db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	users := make([]user, 0)
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserStorage) getUserByID(id string) (user, error) {
	var target user
	row := u.db.QueryRow("select * from users where ID = ?", id)
	err := row.Scan(&target.ID, &target.Name, &target.Email)
	if err != nil {
		panic(err)
	}
	return target, nil
}

package users

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (u *UserStorage) geAllUsers() ([]user, error) {
	rows, err := u.db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	users := make([]user, 0)
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserStorage) getUserByID(id int) (user, error) {
	var target user
	row := u.db.QueryRow("select * from users where ID = ?", id)
	err := row.Scan(&target.ID, &target.Name, &target.Email)
	if err != nil {
		return user{}, err
	}
	return target, nil
}

func (u *UserStorage) createUser(name string, email string) error {
	_, err := u.db.Exec("insert into users (Name, Email) values(?, ?)", name, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStorage) deleteUser(id int) error {
	_, err := u.db.Exec("delete from users where ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}

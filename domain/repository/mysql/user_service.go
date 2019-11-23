package mysql

import (
	"database/sql"
	"fmt"

	"github.com/heroku/go-getting-started/domain/model"

	"github.com/heroku/go-getting-started/domain/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:15
**/

type userService struct {
	db *sql.DB
}

func (us userService) DBConn() *sql.DB {
	return us.db
}

func (us userService) Save(u *model.User, tx *sql.Tx) error {
	query := "INSERT INTO users (id, username, name, password, role, last_login) " +
		" VALUES (?, ?, ?, ?, ?, ?) "
	st, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = st.Exec(u.ID, u.Username, u.Name, u.Password, u.Role, u.LastLogin)
	if err != nil {
		return err
	}
	defer st.Close()
	return err
}

func (us userService) Update(u *model.User, tx *sql.Tx) error {
	query := "UPDATE users " +
		" SET username = ?, " +
		"	name = ?, " +
		" 	password = ?, " +
		"   role = ? " +
		"	deleted_at = ?" +
		" WHERE id = ?"
	st, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = st.Exec(u.Username, u.Name, u.Password, u.Role, u.DeletedAt, u.ID)
	if err != nil {
		return err
	}
	defer st.Close()
	return err
	panic("implement me")
}

func (us userService) Find(id string) (error, *model.User) {
	query := "SELECT id, name, username, password, role, last_login, " +
		" created_at, updated_at " +
		"	FROM users " +
		" WHERE id = ?"
	row := us.db.QueryRow(query, id)
	u := new(model.User)
	err := row.Scan(&u.ID, &u.Name, &u.Username, &u.Password, &u.Role,
		&u.LastLogin, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return err, nil
	}
	return nil, u
}

func (us userService) FindByUsername(username string) (error, *model.User) {
	query := "SELECT id, name, username, password, role, last_login, " +
		" created_at, updated_at " +
		"	FROM users " +
		" WHERE username = ?"
	row := us.db.QueryRow(query, username)
	u := new(model.User)
	err := row.Scan(&u.ID, &u.Name, &u.Username, &u.Password, &u.Role,
		&u.LastLogin, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return err, nil
	}
	return nil, u
}

func (us userService) FindAll(q string, page, size int) (error, []*model.User) {
	query := "SELECT id, name, username, password, role, last_login, " +
		" created_at, updated_at " +
		"	FROM users " +
		" WHERE (name LIKE '%" + q + "%' " +
		" 	OR username LIKE '%" + q + "%') " +
		" LIMIT ? OFFSET ?"
	rows, err := us.db.Query(query, size, (page-1)*size)
	if err != nil {
		return err, nil
	}
	var users []*model.User
	for rows.Next() {
		u := new(model.User)
		err := rows.Scan(&u.ID, &u.Name, &u.Username, &u.Password, &u.Role,
			&u.LastLogin, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	return nil, users
}

func (us userService) FindAllCount(q string) int {
	query := "SELECT count(id) " +
		"	FROM users " +
		" WHERE (name LIKE '%" + q + "%' " +
		" 	OR username LIKE '%" + q + "%') "
	row := us.db.QueryRow(query)
	c := 0
	err := row.Scan(&c)
	if err != nil {
		return 0
	}
	return c
}

func NewUserRepo(db *sql.DB) repository.UserRepository {
	return &userService{db}
}

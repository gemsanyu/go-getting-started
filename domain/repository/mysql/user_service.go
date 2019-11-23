package mysql

import (
	"database/sql"
	"fmt"

	"github.com/mmuflih/go-di-arch/domain/model"

	"github.com/mmuflih/go-di-arch/domain/repository"
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
	query := "INSERT INTO users (id, email, name, phone, password, role, last_login) " +
		" VALUES (?, ?, ?, ?, ?, ?, ?) "
	st, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = st.Exec(u.ID, u.Email, u.Name, u.Phone, u.Password, u.Role, u.LastLogin)
	if err != nil {
		return err
	}
	defer st.Close()
	return err
}

func (us userService) Update(u *model.User, tx *sql.Tx) error {
	query := "UPDATE users " +
		" SET email = ?, " +
		"	name = ?, " +
		"	phone = ?, " +
		" 	password = ?, " +
		"   role = ? " +
		" WHERE id = ?"
	st, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = st.Exec(u.Email, u.Name, u.Phone, u.Password, u.Role, u.ID)
	if err != nil {
		return err
	}
	defer st.Close()
	return err
	panic("implement me")
}

func (us userService) Find(id string) (error, *model.User) {
	query := "SELECT id, name, email, phone, password, role, last_login, " +
		" created_at, updated_at " +
		"	FROM users " +
		" WHERE id = ?"
	row := us.db.QueryRow(query, id)
	u := new(model.User)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Password, &u.Role,
		&u.LastLogin, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return err, nil
	}
	return nil, u
}

func (us userService) FindByEmail(email string) (error, *model.User) {
	query := "SELECT id, name, email, phone, password, role, last_login, " +
		" created_at, updated_at " +
		"	FROM users " +
		" WHERE email = ?"
	row := us.db.QueryRow(query, email)
	u := new(model.User)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Password, &u.Role,
		&u.LastLogin, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return err, nil
	}
	return nil, u
}

func (us userService) FindAll(q string, page, size int) (error, []*model.User) {
	query := "SELECT id, name, email, phone, password, role, last_login, " +
		" created_at, updated_at " +
		"	FROM users " +
		" WHERE (name LIKE '%" + q + "%' " +
		" 	OR email LIKE '%" + q + "%'" +
		" 	OR phone LIKE '%" + q + "%') " +
		" LIMIT ? OFFSET ?"
	rows, err := us.db.Query(query, size, (page-1)*size)
	if err != nil {
		return err, nil
	}
	var users []*model.User
	for rows.Next() {
		u := new(model.User)
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Password, &u.Role,
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
		" 	OR email LIKE '%" + q + "%'" +
		" 	OR phone LIKE '%" + q + "%')"
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

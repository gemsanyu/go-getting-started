package mysql

import (
	"database/sql"

	"github.com/heroku/go-getting-started/domain/model"

	"github.com/heroku/go-getting-started/domain/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:15
**/

type profileService struct {
	db *sql.DB
}

func (ps profileService) DBConn() *sql.DB {
	return ps.db
}

func (ps profileService) Save(p *model.Profile, tx *sql.Tx) error {
	query := "INSERT INTO profiles (id, user_id, name, birth_date, sex, avatar_url) " +
		" VALUES (?, ?, ?, ?, ?, ?) "
	st, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = st.Exec(p.ID, p.UserID, p.Name, p.BirthDate, p.Sex, p.AvatarURL)
	if err != nil {
		return err
	}
	defer st.Close()
	return err
}

func (ps profileService) Update(p *model.Profile, tx *sql.Tx) error {
	query := "UPDATE profiles " +
		" SET name = ?, " +
		"	birth_date = ?, " +
		" 	sex = ?, " +
		"   avatar_url = ? " +
		" WHERE id = ?"
	st, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = st.Exec(p.Name, p.BirthDate, p.Sex, p.AvatarURL, p.ID)
	if err != nil {
		return err
	}
	defer st.Close()
	return err
}

func (ps profileService) Find(id string) (error, *model.Profile) {
	query := "SELECT id, user_id, name, birth_date, sex, avatar_url " +
		"	FROM profiles " +
		" WHERE id = ?"
	row := ps.db.QueryRow(query, id)
	p := new(model.Profile)
	err := row.Scan(&p.ID, &p.UserID, &p.Name, &p.BirthDate, &p.Sex, &p.AvatarURL)
	if err != nil {
		return err, nil
	}
	return nil, p
}

func (ps profileService) FindByUserID(userID string) (error, *model.Profile) {
	query := "SELECT id, user_id, name, birth_date, sex, avatar_url " +
		"	FROM profiles " +
		" WHERE user_id = ?"
	row := ps.db.QueryRow(query, userID)
	p := new(model.Profile)
	err := row.Scan(&p.ID, &p.UserID, &p.Name, &p.BirthDate, &p.Sex, &p.AvatarURL)
	if err != nil {
		return err, nil
	}
	return nil, p
}

func NewProfileRepo(db *sql.DB) repository.ProfileRepository {
	return &profileService{db}
}

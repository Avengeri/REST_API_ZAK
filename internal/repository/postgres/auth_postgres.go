package postgres

import (
	"Interface_droch_3/internal/model"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Set(user *model.User) error {

	sqlStatement := `
INSERT INTO users(id, name) VALUES ($1,$2)
`
	_, err := r.db.Exec(sqlStatement, user.Id, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) Get(id int64) (*model.User, error) {
	sqlStatement := "SELECT name FROM users WHERE id=$1"

	var user model.User

	err := r.db.QueryRow(sqlStatement, id).Scan(&user.Name)
	if err != nil {
		log.Printf("Не удалось выполнить запрос: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *AuthPostgres) Check(id int64) (bool, error) {
	sqlStatement := "SELECT COUNT(*) FROM users WHERE id=$1"

	var count int
	err := r.db.QueryRow(sqlStatement, id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *AuthPostgres) Delete(id int64) error {
	sqlStatement := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(sqlStatement, id)
	if err != nil {
		log.Printf("Не удалось выполнить запрос: %v", err)
		return err
	}
	return nil
}

func (r *AuthPostgres) GetAllId() []int64 {

	sqlNewTable := `CREATE TABLE IF NOT EXISTS users(
    id INTEGER,
    name VARCHAR
)`

	_, err := r.db.Exec(sqlNewTable)
	if err != nil {
		log.Fatal("Не удалось создать таблицу")
	}

	sqlStatement := "SELECT id FROM users;"

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		log.Printf("Не удалось выполнить запрос: %v", err)
	}

	var ids []int64

	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			log.Printf("Не удалось сканировать результат: %v", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Ошибка при чтении строк: %v", err)
	}
	return ids
}

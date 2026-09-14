package user

import (
	"database/sql"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	)`)
	return &PostgresUserRepository{db: db}
}


func (r *PostgresUserRepository) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	var u User
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *PostgresUserRepository) Create(u User) (*User, error) {
	row := r.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email)
	if err := row.Scan(&u.ID); err != nil {
		return nil, err
	}
	return &u, nil
}


func (r *PostgresUserRepository) Update(id int, u User) (*User, error) {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)
	if err != nil {
		return nil, err
	}
	u.ID = id
	return &u, nil
}

func (r *PostgresUserRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

package postgresql

import (
	"BookStore/internal/models"
	"database/sql"
	"errors"
	// "fmt"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

// CreateUser создаёт новую запись в таблице users
func (u *User) CreateUser(user models.User) (int, error) {
	var newUserID int
	query := `
        INSERT INTO users (username, email, password, phone, created_at)
        VALUES ($1, $2, $3, $4, NOW())
        RETURNING user_id
    `
	err := u.db.QueryRow(query, user.Username, user.Email, user.Password, user.Phone).Scan(&newUserID)
	if err != nil {
		return 0, err
	}
	return newUserID, nil
}

// GetUserByEmail ищет пользователя по email
func (u *User) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `
    SELECT user_id, username, email, password, phone, COALESCE(firstname, '') as firstname, COALESCE(lastname, '') as lastname, role, created_at
    FROM users
    WHERE email = $1
    LIMIT 1
`

	row := u.db.QueryRow(query, email)
	err := row.Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Firstname,
		&user.Lastname,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Пользователь не найден
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

func (u *User) GetUserByID(id int) (models.User, error) {
	var user models.User
	query := `
    SELECT user_id, username, email, password, phone, COALESCE(firstname, '') as firstname, COALESCE(lastname, '') as lastname, role, created_at
    FROM users
    WHERE user_id = $1
    LIMIT 1
`
	err := u.db.QueryRow(query, id).Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Firstname,
		&user.Lastname,
		&user.Role,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return user, errors.New("user not found")
	}
	return user, err
}

func (u *User) UpdateUserProfile(user models.User) error {
	query := `
		UPDATE users SET
			username = $1,
			firstname = $2,
			lastname = $3,
			phone = $4
		WHERE user_id = $5
	`
	_, err := u.db.Exec(query, user.Username, user.Firstname, user.Lastname, user.Phone, user.UserId)
	return err
}

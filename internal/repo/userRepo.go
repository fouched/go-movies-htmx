package repo

import (
	"context"
	"errors"
	"github.com/fouched/go-movies-htmx/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// GetUserByID returns a user by ID
func GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT
	    id, first_name, last_name, email, password, created_at, updated_at
    FROM 
        users 
    WHERE id = $1`

	row := db.QueryRowContext(ctx, query, id)

	var u models.User
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUser updates a user
func UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE users SET  
                 first_name = $1 , 
                 last_name = $2, 
                 email = $3,
                 updated_at = $4
             WHERE id = $5`

	_, err := db.ExecContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		time.Now(),
		u.ID)

	if err != nil {
		return err
	}

	return nil
}

// Authenticate authenticates a user
func Authenticate(email, password string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var id int
	var hashedPassword string

	row := db.QueryRowContext(ctx, `SELECT id, password FROM users WHERE email = $1`, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return 0, "", errors.New("wrong password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

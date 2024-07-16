package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestUserDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	createTableQuery := `
	CREATE TABLE users (
		user_id TEXT PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		firstname TEXT,
		lastname TEXT
	);
	`
	_, err = db.Exec(createTableQuery)
	assert.NoError(t, err)

	return db, func() {
		db.Close()
	}
}

func TestCreateUser(t *testing.T) {
	db, teardown := setupTestUserDB(t)
	defer teardown()

	repo := repository.NewUserRepository(db)
	user := models.NewUser("testuser", "password", "Test", "User")

	err := repo.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", user.Username).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestGetUserByUsername(t *testing.T) {
	db, teardown := setupTestUserDB(t)
	defer teardown()

	repo := repository.NewUserRepository(db)
	user := models.NewUser("testuser", "password", "Test", "User")

	_, err := db.Exec(`
		INSERT INTO users (user_id, username, password, firstname, lastname)
		VALUES (?, ?, ?, ?, ?)`,
		user.UserID, user.Username, user.Password, user.FirstName, user.LastName)
	assert.NoError(t, err)

	result, err := repo.GetUserByUsername(context.Background(), user.Username)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

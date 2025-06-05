package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// 1. Basic Database Types
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// 2. Database Connection
func connectDB() (*sql.DB, error) {
	connStr := "postgres://username:password@localhost:5432/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// 3. Basic CRUD Operations
func createUser(db *sql.DB, user User) error {
	query := `
		INSERT INTO users (name, email, created_at)
		VALUES ($1, $2, $3)
		RETURNING id`
	
	err := db.QueryRow(query, user.Name, user.Email, time.Now()).Scan(&user.ID)
	if err != nil {
		return err
	}
	
	return nil
}

func getUser(db *sql.DB, id int) (User, error) {
	user := User{}
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	
	return user, nil
}

func updateUser(db *sql.DB, user User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2
		WHERE id = $3`
	
	_, err := db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func deleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

// 4. Transactions
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Deduct from sender
	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
	if err != nil {
		return err
	}

	// Add to receiver
	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// 5. Prepared Statements
func prepareUserQueries(db *sql.DB) (*sql.Stmt, *sql.Stmt, error) {
	insertStmt, err := db.Prepare(`
		INSERT INTO users (name, email, created_at)
		VALUES ($1, $2, $3)
		RETURNING id`)
	if err != nil {
		return nil, nil, err
	}

	selectStmt, err := db.Prepare(`
		SELECT id, name, email, created_at
		FROM users
		WHERE id = $1`)
	if err != nil {
		return nil, nil, err
	}

	return insertStmt, selectStmt, nil
}

// 6. Connection Pool
func setupConnectionPool(db *sql.DB) {
	// Set maximum number of open connections
	db.SetMaxOpenConns(25)

	// Set maximum number of idle connections
	db.SetMaxIdleConns(5)

	// Set maximum lifetime of a connection
	db.SetConnMaxLifetime(5 * time.Minute)
}

// 7. Query with Joins
type UserWithPosts struct {
	User
	PostCount int
}

func getUsersWithPostCount(db *sql.DB) ([]UserWithPosts, error) {
	query := `
		SELECT u.id, u.name, u.email, u.created_at, COUNT(p.id) as post_count
		FROM users u
		LEFT JOIN posts p ON u.id = p.user_id
		GROUP BY u.id, u.name, u.email, u.created_at`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []UserWithPosts
	for rows.Next() {
		var u UserWithPosts
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.PostCount)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// 8. Batch Operations
func batchInsertUsers(db *sql.DB, users []User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO users (name, email, created_at)
		VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, user := range users {
		_, err = stmt.Exec(user.Name, user.Email, time.Now())
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// 9. Error Handling
func handleDBError(err error) {
	if err == sql.ErrNoRows {
		log.Println("No rows found")
		return
	}
	if err == sql.ErrConnDone {
		log.Println("Connection is done")
		return
	}
	log.Printf("Database error: %v", err)
}

// 10. Context with Timeout
func getUserWithTimeout(db *sql.DB, id int) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := User{}
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	
	err := db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	
	return user, nil
}

func main() {
	// Connect to database
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup connection pool
	setupConnectionPool(db)

	// Example usage
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Create user
	err = createUser(db, user)
	if err != nil {
		handleDBError(err)
	}

	// Get user
	retrievedUser, err := getUser(db, user.ID)
	if err != nil {
		handleDBError(err)
	}
	fmt.Printf("Retrieved user: %+v\n", retrievedUser)

	// Update user
	retrievedUser.Name = "John Updated"
	err = updateUser(db, retrievedUser)
	if err != nil {
		handleDBError(err)
	}

	// Get users with post count
	usersWithPosts, err := getUsersWithPostCount(db)
	if err != nil {
		handleDBError(err)
	}
	fmt.Printf("Users with posts: %+v\n", usersWithPosts)

	// Batch insert
	users := []User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
	}
	err = batchInsertUsers(db, users)
	if err != nil {
		handleDBError(err)
	}

	// Transaction example
	err = transferMoney(db, 1, 2, 100.0)
	if err != nil {
		handleDBError(err)
	}
} 
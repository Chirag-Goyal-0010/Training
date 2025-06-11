package controllers

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/yourusername/flights-project/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	DB *sql.DB
}

func NewAuthController(db *sql.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Force role to be "customer" for public registration
	user.Role = "customer"

	// Hash password
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Insert user into database
	query := `INSERT INTO users (username, password, role, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := ac.DB.QueryRow(query, user.Username, user.Password, user.Role, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user_id": user.ID})
}

// Add a new admin registration endpoint that requires a secret key
func (ac *AuthController) RegisterAdmin(c *gin.Context) {
	// Check for admin registration key
	adminKey := c.GetHeader("X-Admin-Key")
	if adminKey != os.Getenv("ADMIN_REGISTRATION_KEY") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid admin registration key"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Force role to be "admin" for admin registration
	user.Role = "admin"

	// Hash password
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Insert user into database
	query := `INSERT INTO users (username, password, role, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := ac.DB.QueryRow(query, user.Username, user.Password, user.Role, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating admin user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin user created successfully", "user_id": user.ID})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from database
	var user models.User
	query := `SELECT id, username, password, role FROM users WHERE username = $1`
	err := ac.DB.QueryRow(query, loginData.Username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if err := user.CheckPassword(loginData.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

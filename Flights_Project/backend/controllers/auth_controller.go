package controllers

import (
	"net/http"

	"flights-project/middleware"
	"flights-project/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthController handles authentication-related requests
type AuthController struct {
	db *gorm.DB
}

// NewAuthController creates a new auth controller
func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{db: db}
}

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Register handles user registration
func (ac *AuthController) Register(c *gin.Context) {
	var request RegisterRequest
	if !middleware.GetValidatedModel(c, &request) {
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := ac.db.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
		c.Error(middleware.BadRequestError("Email already registered", nil))
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(middleware.InternalServerError("Failed to hash password", err))
		return
	}

	// Create the user
	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
		Role:     "user", // Default role
	}

	if err := ac.db.Create(&user).Error; err != nil {
		c.Error(middleware.InternalServerError("Failed to create user", err))
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(user.ID, user.Email, user.Role, user.Username)
	if err != nil {
		c.Error(middleware.InternalServerError("Failed to generate token", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login handles user login
func (ac *AuthController) Login(c *gin.Context) {
	var request LoginRequest
	if !middleware.GetValidatedModel(c, &request) {
		return
	}

	// Find the user
	var user models.User
	if err := ac.db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.Error(middleware.UnauthorizedError("Invalid email or password", nil))
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Error(middleware.UnauthorizedError("Invalid email or password", nil))
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(user.ID, user.Email, user.Role, user.Username)
	if err != nil {
		c.Error(middleware.InternalServerError("Failed to generate token", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}

// GetProfile returns the current user's profile
func (ac *AuthController) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	role, _ := c.Get("role")
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": gin.H{
			"id":       userID,
			"username": username,
			"email":    email,
			"role":     role,
		},
	})
}

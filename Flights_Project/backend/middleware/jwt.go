package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"flights-project/logger"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// claims defines the structure of our JWT claims
type Claims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// generateToken generates a new JWT token
func GenerateToken(userID uint, email, role, username string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET environment variable not set")
	}

	expirationTime := time.Now().Add(24 * time.Hour) // Default expiration
	if expStr := os.Getenv("JWT_EXPIRATION"); expStr != "" {
		if exp, err := time.ParseDuration(expStr); err == nil {
			expirationTime = time.Now().Add(exp)
		}
	}

	claims := &Claims{
		UserID:   userID,
		Email:    email,
		Role:     role,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.Error("Failed to sign token", zap.Error(err))
		return "", InternalServerError("Failed to create token", err)
	}
	return tokenString, nil
}

// GenerateRefreshToken generates a new refresh token
func GenerateRefreshToken(userID uint) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET environment variable not set")
	}

	expirationTime := time.Now().Add(7 * 24 * time.Hour) // Default refresh expiration
	if expStr := os.Getenv("JWT_REFRESH_EXPIRATION"); expStr != "" {
		if exp, err := time.ParseDuration(expStr); err == nil {
			expirationTime = time.Now().Add(exp)
		}
	}

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.Error("Failed to sign refresh token", zap.Error(err))
		return "", InternalServerError("Failed to create refresh token", err)
	}
	return tokenString, nil
}

// JWTMiddleware validates the JWT token from the request header
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			UnauthorizedError("Authorization header required", nil)
			c.Abort()
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			BadRequestError("Invalid token format: must be 'Bearer <token>'", nil)
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			InternalServerError("JWT_SECRET environment variable not set", nil)
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					UnauthorizedError("Token is expired", err)
				} else {
					UnauthorizedError("Invalid token", err)
				}
			} else {
				InternalServerError("Error parsing token", err)
			}
			c.Abort()
			return
		}

		if !token.Valid {
			UnauthorizedError("Invalid token", nil)
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("username", claims.Username)

		c.Next()
	}
}

// AdminMiddleware checks if the authenticated user has an admin role
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			ForbiddenError("Access denied: Admin role required", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

// RefreshToken handles refreshing an expired access token using a refresh token
func RefreshToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		UnauthorizedError("Authorization header required", nil)
		return
	}

	if !strings.HasPrefix(tokenString, "Bearer ") {
		BadRequestError("Invalid token format: must be 'Bearer <token>'", nil)
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		InternalServerError("JWT_SECRET environment variable not set", nil)
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired == 0 { // Token must be expired for refresh
				UnauthorizedError("Token is not expired, cannot refresh", nil)
				return
			}
		} else {
			InternalServerError("Error parsing refresh token", err)
			return
		}
	}

	if !token.Valid && claims.ExpiresAt < time.Now().Unix() { // Only allow refresh for expired but otherwise valid tokens
		InternalServerError("Invalid refresh token", nil)
		return
	}

	// Generate a new access token
	newToken, err := GenerateToken(claims.UserID, claims.Email, claims.Role, claims.Username)
	if err != nil {
		InternalServerError("Failed to generate new access token", err)
		return
	}

	// Optionally, generate a new refresh token as well, and invalidate the old one
	// For simplicity, we'll just issue a new access token here

	c.JSON(http.StatusOK, gin.H{"access_token": newToken})
}

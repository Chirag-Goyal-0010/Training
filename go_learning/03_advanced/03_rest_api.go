package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 1. Basic Types
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// 2. In-memory Database
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", CreatedAt: time.Now()},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", CreatedAt: time.Now()},
}

// 3. Middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			respondWithError(w, http.StatusUnauthorized, "API key required")
			return
		}
		// In a real application, validate the API key
		next.ServeHTTP(w, r)
	}
}

// 4. Response Helpers
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshaling JSON"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}

// 5. Handlers
func getUsers(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	for _, user := range users {
		if user.ID == userID {
			respondWithJSON(w, http.StatusOK, user)
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "User not found")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user.ID = len(users) + 1
	user.CreatedAt = time.Now()
	users = append(users, user)

	respondWithJSON(w, http.StatusCreated, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	for i, user := range users {
		if user.ID == userID {
			updatedUser.ID = userID
			updatedUser.CreatedAt = user.CreatedAt
			users[i] = updatedUser
			respondWithJSON(w, http.StatusOK, updatedUser)
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "User not found")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			respondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted"})
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "User not found")
}

// 6. File Server
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("static"))
	fs.ServeHTTP(w, r)
}

// 7. CORS Middleware
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func main() {
	// 1. Basic Routes
	http.HandleFunc("/api/users", corsMiddleware(loggingMiddleware(authMiddleware(getUsers))))
	http.HandleFunc("/api/user", corsMiddleware(loggingMiddleware(authMiddleware(getUser))))
	http.HandleFunc("/api/user/create", corsMiddleware(loggingMiddleware(authMiddleware(createUser))))
	http.HandleFunc("/api/user/update", corsMiddleware(loggingMiddleware(authMiddleware(updateUser))))
	http.HandleFunc("/api/user/delete", corsMiddleware(loggingMiddleware(authMiddleware(deleteUser))))

	// 2. Static File Server
	http.HandleFunc("/static/", corsMiddleware(loggingMiddleware(serveStaticFiles)))

	// 3. Start Server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Example API Usage:
/*
GET /api/users - Get all users
GET /api/user?id=1 - Get user by ID
POST /api/user/create - Create new user
PUT /api/user/update?id=1 - Update user
DELETE /api/user/delete?id=1 - Delete user

Headers required:
X-API-Key: your-api-key
Content-Type: application/json

Example POST/PUT body:
{
    "name": "John Doe",
    "email": "john@example.com"
}
*/ 
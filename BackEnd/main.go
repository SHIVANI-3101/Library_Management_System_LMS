package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/jmoiron/sqlx"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-contrib/cors"
)

var db *sql.DB
var jwtKey = []byte("your_secret_key")

func main() {
	// Open the database (create it if it doesn't exist)
	var err error
	db, err = sql.Open("sqlite3", "./taskmanager")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Connection to SQLite database established successfully!")

	// Handling routes of requests
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"} // Allow the Authorization and Content-Type headers
	router.Use(cors.New(config))

	//CREATE USER
	router.POST("/user/create", addUser)
	router.POST("/user/login", loginUser)

	// Protected routes (require authentication)
	admin := router.Group("/admin")
	admin.Use(authMiddleware("Admin"))
	{
		admin.POST("/book/create", addBook)
		admin.POST("/book/update", updateBook)
		admin.GET("/book/delete/:id", removeBook)
		admin.GET("/specific-book/:id", getBookData)
		admin.POST("/user/request/create", requestEvent)
		admin.POST("/request/update", updateUser)
		// admin.POST("/request/delete", deleteUser)
	}

	owner := router.Group("/owner")
	owner.Use(authMiddleware("Creator"))
	{
		owner.POST("/library/create", createLibrary)
		owner.POST("/library/update", updateLibrary)
		owner.GET("/library/delete/:id", removeLibrary)
		owner.GET("/libraries", getAllLibrary)
		owner.GET("/library/search/:query", searchLibrary)
		owner.GET("/library/admin/:id", getAdminData)
		owner.POST("/library/admin/create", addUser)
		owner.POST("/library/admin/update", updateUser)
		owner.GET("/library/:id", getLibraryData)
	}

	users := router.Group("/users")
	users.Use(authMiddleware("Reader"))
	{
		users.GET("/book/search/:query", searchBook)
		users.GET("/books", getAllBooks)
		users.POST("/raiseissue", raiseIssueRequest)
	}

	router.Run("localhost:8080")
}

// Login The User
func loginUser(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the JSON request body to the loginRequest struct
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the database connection
	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Query the database for the user's password and role
	var user Users
	err = db.Get(&user, "SELECT Email,Password, Role FROM Users WHERE Email = ?", loginRequest.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Please enter a registered email"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Compare the stored password with the provided password
	if *user.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please enter the correct password"})
		return
	}

	tokenString, err := generateJWTToken(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the token in a cookie
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "role": user.Role, "token": tokenString})
}

// Generate JWT token
func generateJWTToken(email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"role":  role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// JWT authentication middleware with role check
func authMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Extract the token from the Bearer token format
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}
		tokenString = parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Check if the user has the required role
		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		// Token is valid and user has the required role, continue with the request
		c.Next()
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-contrib/cors"
)

var db *sql.DB

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
	router.Use(cors.New(config))

	//CREATE LIBRARY
	router.POST("/owner/library/create", createLibrary)
	// router.POST("admin/library/update", addUser)
	router.DELETE("owner/library/delete", deleteLibrary)
	// router.POST("admin/library/search", addUser)

	//CREATE USER
	router.POST("/owner/create", addUser)
	// router.POST("/owner/update", addUser)
	router.DELETE("/owner/delete", deleteUser)

	//BOOKS
	router.POST("admin/book/create", addBook)
	router.POST("admin/book/update", updateBook)
	router.GET("admin/book/delete/:id", removeBook)
	router.GET("user/book/search", searchBook)
	router.GET("user/books", getAllBooks)
	router.GET("admin/specific-book/:id", getBookData)

	//REQUESTS CRUD
	router.POST("user/request/create", requestEvent)
	// router.POST("admin/request/update", addUser)
	// router.POST("admin/request/delete", addUser)
	// router.POST("admin/request/delete", addUser)

	//RAISE ISSUE
	router.POST("/user/raiseissue", raiseIssueRequest)

	// //......OWNER.....
	// router.POST("/addOwner", createLibrary)
	// router.POST("/user", addUser)

	// //....ADMIN....
	// router.POST("/addbook", addBook)
	// router.DELETE("/removebook", removeBook)
	// //router.PATCH("/updatebook/:ISBN", updateBook)

	// //.......READER...........
	// //router.GET("/searchbook", searchBook)

	router.Run("localhost:8080")

}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-contrib/cors"
)

var db *sql.DB

// CREATE LIBRARY API
type Library struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatorID int    `json:"creator_id"`
}

func createLibrary(c *gin.Context) {
	var library Library
	if err := c.ShouldBindJSON(&library); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Library (ID, Name, CreatorID) VALUES (?,?,?)", library.ID, library.Name, library.CreatorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Library created successfully"})
}

// DELETE LIBRARY API
func deleteLibrary(c *gin.Context) {
	var library Library
	if err := c.ShouldBindJSON(&library); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Library WHERE ID = ?", library.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Library deleted successfully"})
}

//ADD USER API

type Users struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	ContactNumber int    `json:"contact_number"`
	Role          string `json:"role"`
	LibID         int    `json:"lib_id"`
}

func addUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Users (ID, Name, Email, ContactNumber, Role, LibID) VALUES (?,?,?,?,?,?)", user.ID, user.Name, user.Email, user.ContactNumber, user.Role, user.LibID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

//DELETE USER API

func deleteUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Users WHERE ID = ?", user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

//...........ADMIN...........
//CREATE BOOK API

type BookInventory struct {
	ISBN            int    `json:"isbn"`
	LibID           int    `json:"lib_id"`
	Title           string `json:"title"`
	Authors         string `json:"authors"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
}

func addBook(c *gin.Context) {
	var book BookInventory
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.Exec("INSERT INTO BookInventory (ISBN, LibID, Title, Authors, Publisher, Version, TotalCopies, AvailableCopies) VALUES (?,?,?,?,?,?,?,?)", book.ISBN, book.LibID, book.Title, book.Authors, book.Publisher, book.Version, book.TotalCopies, book.AvailableCopies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
}

//Remove BOOK API

func removeBook(c *gin.Context) {
	var book BookInventory
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.Exec("DELETE FROM BookInventory WHERE ISBN = ?", book.ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book removed successfully"})
}

// UPDATE BOOK API
/*
func updateBook(c *gin.Context) {
	var book BookInventory

		if bookS.ISBN =""{
			c.JSON(http.StatusBadRequest, gin.H{"message: id Required"})
			return
		}
		err:= Lib.DB.Where("isbn= ?", isbn).First(&book).Error
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"message: Book with this isbn does not exist"})
			return
		}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.Exec("UPDATE BookInventory SET Title=?, Authors=?, Publisher=?, Version=?, TotalCopies=?, AvailableCopies = ? WHERE ISBN = ?", book.ISBN, book.LibID, book.Title, book.Authors, book.Publisher, book.Version, book.TotalCopies, book.AvailableCopies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book Updted successfully"})
}
*/

// SEARCH BOOK API

func searchBook(c *gin.Context) {
	var book BookInventory
	if err := c.ShouldBindQuery(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Queryx("SELECT * FROM BookInventory WHERE Title = ?", book.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []BookInventory
	for rows.Next() {
		var b BookInventory
		err := rows.StructScan(&b)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, b)
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

// Request Book api

type RequestEvents struct {
	ReqID        int    `json:"req_id"`
	BookID       int    `json:"book_id"`
	ReaderID     int    `json:"reader_id"`
	RequestDate  string `json:"request_date"`
	ApprovalDate string `json:"approval_date"`
	ApproverID   int    `json:"approver_id"`
	RequestType  string `json:"request_type"`
}

func requestEvent(c *gin.Context) {
	var event RequestEvents
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO RequestEvents (ReqID, BookID, ReaderID, RequestDate, ApprovalDate, ApproverID, RequestType) VALUES (?,?,?,?,?,?,?)", event.ReqID, event.BookID, event.ReaderID, event.RequestDate, event.ApprovalDate, event.ApproverID, event.RequestType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request done successfully"})
}

//RAISE ISSUE REQUEST API

type IssueRegistry struct {
	IssueID            int    `json:"issue_id"`
	ISBN               int    `json:"isbn"`
	ReaderID           int    `json:"reader_id"`
	IssueApproverID    int    `json:"issue_approver_id"`
	IssueStatus        string `json:"issue_status"`
	IssueDate          string `json:"issue_date"`
	ExpectedReturnDate string `json:"expected_return_date"`
	ReturnDate         string `json:"return_date"`
	ReturnApproverID   int    `json:"return_approver_id"`
}

func raiseIssueRequest(c *gin.Context) {
	var request IssueRegistry
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO IssueRegistry (IssueID, ISBN, ReaderID, IssueApproverID, IssueStatus, IssueDate, ExpectedReturnDate, ReturnDate, ReturnApproverID) VALUES (?,?,?,?,?,?,?,?,?)", request.IssueID, request.ISBN, request.ReaderID, request.IssueApproverID, request.IssueStatus, request.IssueDate, request.ExpectedReturnDate, request.ReturnDate, request.ReturnApproverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue request raised successfully"})
}

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

	//BOOKS CRUD
	router.POST("admin/book/create", addBook)
	// router.POST("admin/book/update", addUser)
	router.DELETE("admin/book/delete", removeBook)
	router.GET("user/book/search", searchBook)

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

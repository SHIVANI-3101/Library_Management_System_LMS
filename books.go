package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//...........ADMIN...........
//CREATE BOOK API

type BookInventory struct {
	ID              int    `db:"ID" json:"id"`
	ISBN            int    `db:"ISBN" json:"isbn"`
	LibID           int    `db:"LibID" json:"lib_id"`
	Title           string `db:"Title" json:"title"`
	Authors         string `db:"Authors" json:"authors"`
	Publisher       string `db:"Publisher" json:"publisher"`
	Version         string `db:"Version" json:"version"`
	TotalCopies     int    `db:"TotalCopies" json:"total_copies"`
	AvailableCopies int    `db:"AvailableCopies" json:"available_copies"`
}

func getAllBooks(c *gin.Context) {
	var books []BookInventory

	// Connect to the database
	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Execute the query to fetch all books
	if err := db.Select(&books, "SELECT * FROM BookInventory"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the fetched data as JSON response
	c.JSON(http.StatusOK, books)
}

func getBookData(c *gin.Context) {
	// Get the book ID from the request parameters
	bookID := c.Param("id")

	// Connect to the database
	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Execute the query to fetch the data of the specific book
	var book BookInventory
	if err := db.Get(&book, "SELECT * FROM BookInventory WHERE ID = ?", bookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the fetched data as JSON response
	c.JSON(http.StatusOK, book)
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

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func updateBook(c *gin.Context) {
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
	defer db.Close()

	_, err = db.Exec("UPDATE BookInventory SET ISBN=?, LibID=?, Title=?, Authors=?, Publisher=?, Version=?, TotalCopies=?, AvailableCopies=? WHERE ID=?", book.ISBN, book.LibID, book.Title, book.Authors, book.Publisher, book.Version, book.TotalCopies, book.AvailableCopies, book.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func removeBook(c *gin.Context) {
	bookID := c.Param("id")

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.Exec("DELETE FROM BookInventory WHERE ID = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

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

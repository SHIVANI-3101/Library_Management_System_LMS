package main

import (
	"fmt"
	"net/http"

	// "bytes"
	// "image/png"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/skip2/go-qrcode"
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

	// Generate QR code from book data
	qrData := fmt.Sprintf("Title: %s\nAuthors: %s\nPublisher: %s\nISBN: %s", book.Title, book.Authors, book.Publisher, strconv.Itoa(book.ISBN))
	qrCode, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	// Save QR code image to file
	qrFileName := fmt.Sprintf("%s.png", book.Title)
	err = saveQRCodeToFile(qrCode, qrFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save QR code to file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func saveQRCodeToFile(qrCode []byte, fileName string) error {
	// Prepend the folder name to the file name
	fileNameWithPath := filepath.Join("qrcodes", fileName)

	file, err := os.Create(fileNameWithPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(qrCode)
	if err != nil {
		return err
	}
	return nil
}

func deleteQRCodeFile(fileName string) error {
	filePath := filepath.Join("qrcodes", fileName)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
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

	err = deleteQRCodeFile(fmt.Sprintf("%s.png", book.Title))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete QR code file"})
		return
	}

	// Generate QR code from book data
	qrData := fmt.Sprintf("Title: %s\nAuthors: %s\nPublisher: %s\nISBN: %s", book.Title, book.Authors, book.Publisher, strconv.Itoa(book.ISBN))
	qrCode, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	// Save QR code image to file
	qrFileName := fmt.Sprintf("%s.png", book.Title)
	err = saveQRCodeToFile(qrCode, qrFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save QR code to file"})
		return
	}

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

	// Retrieve book details to get the title for QR code deletion
	var book BookInventory
	err = db.Get(&book, "SELECT Title FROM BookInventory WHERE ID = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete QR code file
	err = deleteQRCodeFile(fmt.Sprintf("%s.png", book.Title))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete QR code file"})
		return
	}

	// Delete book from database
	db.Exec("DELETE FROM BookInventory WHERE ID = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// SEARCH BOOK API

func searchBook(c *gin.Context) {
	query := c.Param("query")

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Queryx("SELECT * FROM BookInventory WHERE Title LIKE '%' || ? || '%' OR Authors LIKE '%' || ? || '%' OR Publisher LIKE '%' || ? || '%'", query, query, query)
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

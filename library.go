package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

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

	_, err = db.Exec("INSERT INTO Library (Name, CreatorID) VALUES (?,?)", library.Name, library.CreatorID)
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

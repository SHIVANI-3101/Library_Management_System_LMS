package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// CREATE LIBRARY API
type Library struct {
	ID        int    `db:"ID" json:"id"`
	Name      string `db:"Name" json:"name"`
	CreatorID *int   `db:"CreatorID" json:"creator_id"`
}

func getAllLibrary(c *gin.Context) {

	var libraries []Library

	// Connect to the database
	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Execute the query to fetch the data of the specific library
	if err := db.Select(&libraries, "SELECT * FROM Library"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the fetched data as JSON response
	c.JSON(http.StatusOK, libraries)
}

func getLibraryData(c *gin.Context) {
	// Get the library ID from the request parameters
	libraryID := c.Param("id")

	// Connect to the database
	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Execute the query to fetch the data of the specific library
	var library Library
	if err := db.Get(&library, "SELECT * FROM Library WHERE ID = ?", libraryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the fetched data as JSON response
	c.JSON(http.StatusOK, library)
}

func searchLibrary(c *gin.Context) {
	query := c.Param("query")

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var libraries []Library
	err = db.Select(&libraries, "SELECT * FROM Library WHERE Name LIKE '%' || ? || '%'", query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"libraries": libraries})
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

	result, err := db.Exec("INSERT INTO Library (Name, CreatorID) VALUES (?,?)", library.Name, library.CreatorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	libraryID, _ := result.LastInsertId()

	c.JSON(http.StatusOK, gin.H{"message": "success", "library_id": libraryID})
}

// DELETE LIBRARY API
func removeLibrary(c *gin.Context) {

	libraryID := c.Param("id")

	db, err := sqlx.Connect("sqlite3", "Lib.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Library WHERE ID = ?", libraryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func updateLibrary(c *gin.Context) {
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

	_, err = db.Exec("UPDATE Library SET Name = ? WHERE ID = ?", library.Name, library.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

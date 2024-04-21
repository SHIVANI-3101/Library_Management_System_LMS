package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//ADD USER API

type Users struct {
	ID            int     `db:"ID" json:"id"`
	Name          string  `db:"Name" json:"name"`
	Email         string  `db:"Email" json:"email"`
	ContactNumber int     `db:"ContactNumber" json:"contact_number"`
	Role          string  `db:"Role" json:"role"`
	LibID         int     `db:"LibID" json:"lib_id"`
	Password      *string `db:"PASSWORD" json:"pass"`
}

func getAdminData(c *gin.Context) {
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
	var user Users
	if err := db.Get(&user, "SELECT * FROM Users WHERE LibID = ? AND Role='Admin'", libraryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the fetched data as JSON response
	c.JSON(http.StatusOK, user)
}

func setAdminData(c *gin.Context)
{
	
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

	_, err = db.Exec("INSERT INTO Users (Name, Email, ContactNumber, Role, LibID) VALUES (?,?,?,?,?)", user.ID, user.Name, user.Email, user.ContactNumber, user.Role, user.LibID)
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

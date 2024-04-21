package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

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

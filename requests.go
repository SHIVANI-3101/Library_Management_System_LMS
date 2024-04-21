package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

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

// // DATABASE CONNECTIVITY AND CREATE SCHEMA
// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {

// 	db, err := sql.Open("sqlite3", "Lib.db")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()
// 	// Create the database schema
// 	createTables := `
//     CREATE TABLE IF NOT EXISTS Library (
//         ID INTEGER PRIMARY KEY AUTOINCREMENT,
//         Name VARCHAR(255) UNIQUE,
//         CreatorID INT,
//         FOREIGN KEY (CreatorID) REFERENCES Users(ID)
//     );

//     CREATE TABLE IF NOT EXISTS Users (
//         ID INTEGER PRIMARY KEY AUTOINCREMENT,
//         Name VARCHAR(255),
//         Email VARCHAR(255) UNIQUE,
//         ContactNumber VARCHAR(20),
//         Role VARCHAR(10),
//         LibID INT,
//         FOREIGN KEY (LibID) REFERENCES Library(ID)
//     );

//     CREATE TABLE IF NOT EXISTS BookInventory (
//         ID INTEGER PRIMARY KEY AUTOINCREMENT,
//         ISBN INT UNIQUE,
//         LibID INT,
//         Title VARCHAR(255),
//         Authors TEXT,
//         Publisher VARCHAR(255),
//         Version VARCHAR(50),
//         TotalCopies INT,
//         AvailableCopies INT,
//         FOREIGN KEY (LibID) REFERENCES Library(ID)
//     );

//     CREATE TABLE IF NOT EXISTS RequestEvents (
//         ReqID INTEGER PRIMARY KEY AUTOINCREMENT,
//         BookID VARCHAR(20),
//         ReaderID INT,
//         RequestDate TIMESTAMP,
//         ApprovalDate TIMESTAMP,
//         ApproverID INT,
//         RequestType VARCHAR(10),
//         FOREIGN KEY (BookID) REFERENCES BookInventory(ISBN),
//         FOREIGN KEY (ReaderID) REFERENCES Users(ID),
//         FOREIGN KEY (ApproverID) REFERENCES Users(ID)
//     );

//     CREATE TABLE IF NOT EXISTS IssueRegistry (
//         IssueID INTEGER PRIMARY KEY AUTOINCREMENT,
//         ISBN INT,
//         ReaderID INT,
//         IssueApproverID INT,
//         IssueStatus VARCHAR(20),
//         IssueDate TIMESTAMP,
//         ExpectedReturnDate TIMESTAMP,
//         ReturnDate TIMESTAMP,
//         ReturnApproverID INT,
//         FOREIGN KEY (ISBN) REFERENCES BookInventory(ISBN),
//         FOREIGN KEY (ReaderID) REFERENCES Users(ID),
//         FOREIGN KEY (IssueApproverID) REFERENCES Users(ID),
//         FOREIGN KEY (ReturnApproverID) REFERENCES Users(ID)
//     );
//     `

// 	_, err = db.Exec(createTables)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Database schema created successfully")

// }
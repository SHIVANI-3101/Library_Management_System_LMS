# Library Management System

## Overview
The **Library Management System** is a web-based application designed to manage a library efficiently. It allows administrators to perform full CRUD (Create, Read, Update, Delete) operations, while users have retrieval capabilities along with limited CRUD operations. This system streamlines the process of managing books, users, and borrowing records.

## Features
### Admin Features
- Create, update, delete, and view books.
- Manage users (add, update, delete accounts).
- Track borrowing and return records.
- Manage library details and settings.

### User Features
- Search and retrieve book details.
- Borrow and return books.
- Update user profile.
- View borrowing history.

## Technologies Used
- **Frontend:** HTML, CSS, JavaScript
- **Backend:** Go
- **Database:** SQL (MySQL/PostgreSQL/SQLite)

## SQL Operations
This system utilizes various SQL operations, including:
- **Creating tables** for books, users, and borrowing records.
- **Inserting records** when adding new books and users.
- **Retrieving data** to fetch available books and user details.
- **Updating records** to modify book availability and user profiles.
- **Deleting records** when removing books or users from the system.

## Installation & Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/library-management.git
   ```
2. Navigate to the project directory:
   ```sh
   cd library-management
   ```
3. Set up the database:
   ```sh
   go run setup_db.go
   ```
4. Start the backend server:
   ```sh
   go run main.go
   ```
5. Open `index.html` in a browser to access the frontend.

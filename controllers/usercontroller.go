package controllers

import (
	"fmt"
	m "latFramework/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// buat print semua data user yang ada di tabel database
func GetUser(c *gin.Context) {
	db := connect()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println("Error querying database:", err)
		c.JSON(500, gin.H{"error": "internal server error 1 "})
		return
	}
	defer rows.Close()

	var users []m.Users
	for rows.Next() {
		var user m.Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			log.Println("Error scanning row:", err)
			c.JSON(500, gin.H{"error": "internal server error 2"})
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		c.JSON(500, gin.H{"error": "internal server error 3"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// buat update data user by id
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	var user m.Users
	// inputan datanya pake raw JSON
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}

	// Check nama sama Age yg di input di JSON
	if user.Name == "" || user.Age <= 0 {
		c.JSON(400, gin.H{"error": "name and age are required"})
		return
	}

	// connect dan close db
	db := connect()
	defer db.Close()

	query := "UPDATE users SET name=?, age=? WHERE id=?"

	_, err = db.Exec(query, user.Name, user.Age, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to update user"})
		fmt.Println("Error updating user:", err)
		return
	}

	// Fetch updated user data
	updatedUser := m.Users{ID: userID, Name: user.Name, Age: user.Age}

	c.JSON(200, updatedUser)
}

// buat insert user baru
func InsertNewUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	var user struct {
		Name string `form:"name" binding:"required"`
		Age  int    `form:"age" binding:"required"`
	}

	if err := c.ShouldBind(&user); err != nil {
		log.Println("Error binding form data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: Incomplete data"})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error 1"})
		return
	}
	defer tx.Rollback()

	// Execute SQL query to insert user
	query := "INSERT INTO users (name, age) VALUES (?, ?)"
	result, err := tx.Exec(query, user.Name, user.Age)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error 2"})
		return
	}
	lastInsertID, _ := result.LastInsertId()

	// Commit
	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error 3"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User inserted successfully with ID: %d", lastInsertID)})
}

// buat hapus data user by id
func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer tx.Rollback()

	// Delete query
	query := "DELETE FROM users WHERE id = ?"

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Println("Error preparing SQL statement:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer stmt.Close()

	// Execute the delete query
	_, err = stmt.Exec(userID)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID %d deleted successfully", userID)})
}

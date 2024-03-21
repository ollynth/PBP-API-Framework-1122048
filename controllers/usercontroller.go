package controllers

import (
	m "latFramework/models"
	"log"
	"net/http"

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

	// Iterate over the rows and handle the data
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

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		c.JSON(500, gin.H{"error": "internal server error 3"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// buat insert user baru

// buat update data user by id

// buat hapus data user by id

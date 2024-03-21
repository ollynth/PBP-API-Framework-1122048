package controllers

import (
	m "latFramework/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// buat print semua data user yang ada di tabel database
func GetUser(c *gin.Context) {
	var user []m.Users
	_, err := dbmap.Select(&user, "select * from users")
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

// buat insert user baru

// buat update data user by id

// buat hapus data user by id

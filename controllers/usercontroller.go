package controllers

import (
	m "latFramework/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// buat print semua data user yang ada di tabel database
func GetUser(c *gin.Context) {
	dbmap := GetDBMap()
	var user []m.Users
	query := "SELECT * FROM users"
	_, err := dbmap.Select(&user, query)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// buat insert user baru

// buat update data user by id

// buat hapus data user by id

package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latpbp_gin_framework?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	// Check if the connection is established
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Connected to the database")
	return db
}

// nyobain buat connection pake gin tp masih error
// func connect() *gorp.DbMap {
// 	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db_latpbp_gin_framework")
// 	checkErr(err, "sql.Open failed")
// 	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
// 	// err = dbmap.CreateTablesIfNotExists()
// 	// checkErr(err, "Create tables failed")
// 	return dbmap
// }
// func checkErr(err error, msg string) {
// 	if err != nil {
// 		log.Fatalln(msg, err)
// 	}
// }
// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
// 		c.Next()
// 	}
// }

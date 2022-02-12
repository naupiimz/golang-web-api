package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "db_golang"
)

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)

	r.GET("/product/:id", productHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "root",
	})
}

func productHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

package main

import (
	"database/sql"
	"dateservice/pkg/repo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const (
	dbData = "user=postgres password=1 dbname=postgres sslmode=disable"
)

func main() {
	db, err := sql.Open("postgres", dbData)
	if err != nil {
		log.Println(err)
	}
	log.Println("successful connect to database")
	rep, err := repo.NewRepo(db)
	if err != nil {
		return
	}
	data, err := rep.GetActivityByDate(time.Now())
	router := gin.Default()
	router.GET("/data", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Users": data,
		})
	})
	router.Run(":3001")

}

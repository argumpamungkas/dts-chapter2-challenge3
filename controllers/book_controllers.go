package controllers

// import (
// 	// "chapter2-challenge-sesi-3/models"
// 	"chapter2-challenge-sesi-3/models"
// 	"database/sql"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	db  *sql.DB
// 	err error
// )

// func GetAllBook(ctx *gin.Context) {
// 	var allBooks []models.Book

// 	sqlStatement := `SELECT * FROM books`

// 	rows, err := db.Query(sqlStatement)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var books models.Book

// 		err = rows.Scan(&books.ID, &books.Title, &books.Author, &books.Description)

// 		if err != nil {
// 			log.Println("Error di db query")
// 			return
// 		}

// 		allBooks = append(allBooks, books)
// 	}

// 	if allBooks == nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"data":        []string{},
// 			"status_code": http.StatusOK,
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"books": allBooks,
// 	})

// }

package repo

import (
	"chapter2-challenge-sesi-3/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "gume98"
	dbname   = "db-book-sql"
)

var (
	db  *sql.DB
	err error
)

func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Connect to database")

}

func GetAllBook(ctx *gin.Context) {
	var allBooks []models.Book

	sqlStatement := `SELECT * FROM books ORDER BY id`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var books models.Book

		err = rows.Scan(&books.ID, &books.Title, &books.Author, &books.Description)

		if err != nil {
			log.Println("Error di db query")
			return
		}

		allBooks = append(allBooks, books)
	}

	if allBooks == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data":        []string{},
			"status_code": http.StatusOK,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": allBooks,
	})

}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookData models.Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid convert")
		return
	}

	sqlStatement := `SELECT * FROM books WHERE id = $1`
	err = db.QueryRow(sqlStatement, idBook).Scan(&bookData.ID, &bookData.Title, &bookData.Author, &bookData.Description)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":     fmt.Sprintf("book with id %v not found", idBook),
			"status_code": http.StatusNotFound,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":        bookData,
		"status_code": http.StatusOK,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook models.Book
	var lastID int

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	getLastId := `SELECT id FROM books ORDER BY id DESC LIMIT 1`
	err = db.QueryRow(getLastId).Scan(&lastID)
	if err != nil {
		newBook.ID = 1
	}

	newBook.ID = lastID + 1

	sqlStatement := `INSERT INTO books (id, title, author, description) VALUES ($1, $2, $3, $4) Returning *`
	_, err = db.Exec(sqlStatement, newBook.ID, newBook.Title, newBook.Author, newBook.Description)
	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("book %v is created", newBook.Title),
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var updateBook models.Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid Convert")
		return
	}

	if err = ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// GET ID
	findId := `SELECT id FROM books WHERE id = $1`
	err = db.QueryRow(findId, idBook).Scan(&idBook)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":     fmt.Sprintf("book with id %d not found", idBook),
			"status_code": http.StatusNotFound,
		})
		return
	}

	sqlStatement := `UPDATE books SET title = $2, author = $3, description = $4 WHERE id = $1`
	_, err = db.Exec(sqlStatement, idBook, updateBook.Title, updateBook.Author, updateBook.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("book with id %v invalid updated", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %d successfully updated", idBook),
	})

}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid Convert")
		return
	}

	findId := `SELECT id FROM books WHERE id = $1`
	err = db.QueryRow(findId, idBook).Scan(&idBook)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":     fmt.Sprintf("book with id %d not found", idBook),
			"status_code": http.StatusNotFound,
		})
		return
	}

	sqlStatement := `DELETE FROM books WHERE id = $1`
	_, err = db.Exec(sqlStatement, idBook)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %d successfully deleted", idBook),
	})

}

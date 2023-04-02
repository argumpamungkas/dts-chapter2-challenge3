package controllers

import (
	"chapter2-challenge-sesi-3/models"
	"chapter2-challenge-sesi-3/repo"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	var allBook []models.Book

	books, err := repo.GetAllBookDB(allBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Request")
		return
	}

	if books == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": []string{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookData models.Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid convert")
		return
	}

	book, err := repo.GetBookByIdDB(idBook, bookData)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("book with id %v not found", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := repo.CreateBook(book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid create book data")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New book is created",
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var book models.Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid Convert")
		return
	}

	if err = ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = repo.UpdateBookDB(idBook, book)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("book with id %d not found", idBook),
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

	err = repo.DeleteBookDB(idBook)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("book with id %d not found", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %d successfully deleted", idBook),
	})

}

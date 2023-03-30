package router

import (
	"chapter2-challenge-sesi-3/repo"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", repo.GetAllBook)

	router.GET("/books/:bookID", repo.GetBookById)

	router.POST("/books", repo.CreateBook)

	router.PUT("/books/:bookID", repo.UpdateBook)

	router.DELETE("/books/:bookID", repo.DeleteBook)

	return router
}

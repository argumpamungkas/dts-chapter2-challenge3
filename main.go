package main

import (
	"chapter2-challenge-sesi-3/repo"
	"chapter2-challenge-sesi-3/router"
)

func main() {

	repo.ConnectDatabase()

	router.StartServer().Run(":8080")

}

package main

import (
	"pustaka_api/database"
	"pustaka_api/router"
)

func main() {

	database.GetConnection()
	router.Router()
}

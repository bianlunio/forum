package main

import (
	"forum/models"
	"forum/routers"
)

func main() {
	db := models.Connect()
	defer db.Close()
	r := routers.SetRouter()
	r.Run(":8080")
}

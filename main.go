package main

import (
	"forum/models"
	"forum/routers"
)

func main() {
	defer models.Session.Close()
	r := routers.SetRouter()
	r.Run(":8080")
}

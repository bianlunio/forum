package main

import (
	"github.com/bianlunio/forum/models"
	"github.com/bianlunio/forum/routers"
)

func main() {
	defer models.Session.Close()
	r := routers.SetRouter()
	r.Run(":8080")
}

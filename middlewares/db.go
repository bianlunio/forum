package middlewares

import (
	. "github.com/bianlunio/forum/models"

	"github.com/gin-gonic/gin"
)

func BindDBSession(c *gin.Context) {
	// http://blog.mongodb.org/post/80579086742/running-mongodb-queries-concurrently-with-go
	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	reqSession := Session.Clone()
	defer reqSession.Close()

	dbName := "forum"
	if gin.Mode() == gin.TestMode {
		dbName = "forum_test"
	}

	c.Set("db", reqSession.DB(dbName))
	c.Next() // we need to execute the rest handlers of the chain here because the session is released when we leave this scope
}

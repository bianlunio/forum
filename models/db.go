package models

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var DEBUG = false

var db *pg.DB

func Connect() *pg.DB {
	addr := "localhost:5432"

	if os.Getenv("DRONE") == "true" {
		addr = "forum-db:5432"
	}
	db = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     "forum",
		Password: "TSxdMxWB21Bt4j36",
		Database: "forum",
	})

	if gin.Mode() != gin.ReleaseMode && DEBUG {
		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}

			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}

	temp := false
	if gin.Mode() == gin.TestMode {
		orm.SetTableNameInflector(func(s string) string {
			return "test_" + s
		})
		temp = true
	}
	err := createSchema(db, temp)
	handleDBError(err)
	return db
}

func createSchema(db *pg.DB, temp bool) error {
	models := []interface{}{&Topic{}, &Tab{}}
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
			Temp:        temp,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func ClearAll() {
	models := []interface{}{&Topic{}, &Tab{}}
	for _, model := range models {
		err := db.DropTable(model, nil)
		err = db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		handleDBError(err)
	}
}

func handleDBError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

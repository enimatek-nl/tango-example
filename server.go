package main

import (
	"github.com/enimatek-nl/tango-example/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&server.Todo{}); err != nil {
		panic(err)
	}

	api := server.Api{Db: db}
	http.HandleFunc("/api", api.Index)
	http.HandleFunc("/api/todo", api.Process)
	http.Handle("/", http.FileServer(http.Dir(`./web/static/`)))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

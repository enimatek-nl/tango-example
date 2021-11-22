module github.com/enimatek-nl/tango-example

go 1.17

require (
	github.com/enimatek-nl/tango v0.0.0-20211026051839-8b79cf06c137
	gorm.io/driver/sqlite v1.2.3
	gorm.io/gorm v1.22.2
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
)

replace github.com/enimatek-nl/tango => ../tango

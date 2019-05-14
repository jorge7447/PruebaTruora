package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func Connection() *gorm.DB {

	const addr = "postgresql://jsandoval@localhost:26257/truora?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.jsandoval.key&sslcert=certs/client.jsandoval.crt"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error connecting to the database: ", err)
	}

	//db.AutoMigrate(&Key{})
	return db
}

func CloseConnection(db *gorm.DB) {
	defer db.Close()
}

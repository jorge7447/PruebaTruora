package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func Connection() *gorm.DB {

	/*db, err := sql.Open("postgres",
		"postgresql://jsandoval@localhost:26257/truora?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.jsandoval.key&sslcert=certs/client.jsandoval.crt")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
		fmt.Println("error connecting to the database: ", err)
	}
	defer db.Close()

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal(err)
	}*/

	const addr = "postgresql://jsandoval@localhost:26257/truora?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.jsandoval.key&sslcert=certs/client.jsandoval.crt"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error connecting to the database: ", err)
	}
	//defer db.Close()

	//db.AutoMigrate(&Key{})

	return db
}

func CloseConnection(db *gorm.DB) {
	defer db.Close()
}

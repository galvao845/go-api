package dataBaseInstance

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	//"gorm.io/driver/sqlite"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

const driver = "postgres"
const host = "127.0.0.1"
const port = "5432"
const user = "go"
const password = "go"
const dbName = "go"

var sourceName = fmt.Sprintf("host=%s port=%s user=%s"+" password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
var db *sql.DB
var err error

/*
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
*/

func throwError(err error) {
	if err != nil {
		//panic(err.Error())
		fmt.Println("erro de conexao..")
	} else {
		fmt.Println("Conectado..")
	}
}

func ConnectDb() *sql.DB {
	fmt.Println("Iniciando conexao..")
	db, err := sql.Open(driver, sourceName)
	throwError(err)
	return db
}

/*
		  instrumentedDB, err := xray.SQLContext("postgres", "postgres://postgres:q1Q!q1Q!@127.0.0.1:9832/brcconfig")
		  // Handle err

		  db, err := gorm.Open(postgres.New(postgres.Config{
		    Conn: instrumentedDB,
		  }), &gorm.Config{})

					db, err := gorm.Open(postgres.New(postgres.Config{
						DSN:                  "host=192.168.254.246 user=postgres password=q1Q!q1Q! dbname=brcconfig port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
						PreferSimpleProtocol: true, // disables implicit prepared statement usage
					}), &gorm.Config{})

	//dsn := "host=127.0.0.1 user=postgres password=q1Q!q1Q! dbname=brcconfig port=9832 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Fatal(db)
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		return db
			// Migrate the schema
			db.AutoMigrate(&Product{})

			// Create
			db.Create(&Product{Code: "D42", Price: 100})

			// Read
			var product Product
			db.First(&product, 1)                 // find product with integer primary key
			db.First(&product, "code = ?", "D42") // find product with code D42

			// Update - update product's price to 200
			db.Model(&product).Update("Price", 200)
			// Update - update multiple fields
			db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
			db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

			// Delete - delete product
			db.Delete(&product, 1)*/

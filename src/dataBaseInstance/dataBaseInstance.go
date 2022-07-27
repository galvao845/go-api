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

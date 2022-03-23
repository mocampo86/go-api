package main

import (
	"context"
	"fmt"
	"golang-ms-example/product"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
)

func main() {
	ctx := context.Background()

	//SQL Repository
	//db, err := sql.Open("mssql", config.Values.DB.ConnectionString())
	//repository := product.NewSQLRepository(db)

	//GORM Repository
	db, err := gorm.Open("sqlite3", "./gorm.db")
	repository := product.NewGormRepository(db)
	db.AutoMigrate(&product.DAOProduct{})

	if err != nil {
		fmt.Println("No hay db...")
		panic(err)
	} else {
		fmt.Println("Conexion con DB con exito")
	}

	service := product.NewService(repository)
	endpoints := product.MakeEndpoints(service)

	err = http.ListenAndServe(":8000", product.NewHTTPServer(ctx, endpoints))
	if err != nil {
		panic(err)
	}
}

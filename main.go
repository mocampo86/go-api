package main

import (
	"context"
	"database/sql"
	"fmt"
	"golang-ms-example/config"
	"golang-ms-example/product"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	ctx := context.Background()

	//Repository
	db, err := sql.Open("mssql", config.Values.DB.ConnectionString())

	if err != nil {
		fmt.Println("No hay db...")
		panic(err)
	} else {
		fmt.Println("Conexion con DB con exito")
	}

	repository := product.NewRepository(db)
	service := product.NewService(repository)
	endpoints := product.MakeEndpoints(service)

	err = http.ListenAndServe(":8000", product.NewHTTPServer(ctx, endpoints))
	if err != nil {
		panic(err)
	}
}

package product

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"
)

type sqlrepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Repository {
	return &sqlrepository{
		db: db,
	}
}

func (r sqlrepository) Remove(ctx context.Context, id int64) error {
	query := "DELETE FROM Products WHERE ProductId = ?"
	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		log.Fatal("Could not fetch data from the database with error: ", row.Err())
	}
	return row.Err()
}

//This func is not used any more. Search encapsulates Find by Id and Find by Name
func (r sqlrepository) Find_Obsolete(ctx context.Context, id int64) (product DAOProduct, err error) {
	query := "SELECT ProductId,Name,Description,Price,SKU FROM Products WHERE ProductId = ?"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Fatal("Could not fetch data from the database with error: ", err)
		return DAOProduct{}, errors.New("Something get wrong with data base")
	}

	result, err := scanProductRows(rows)
	if err != nil {
		log.Fatal("Could not scan Rows with error: ", err)
		return DAOProduct{}, errors.New("Could not scan Rows")
	}

	if len(result) > 1 {
		log.Fatal("Too many results")
		return DAOProduct{}, errors.New("Too many results")
	}
	if len(result) == 0 {
		return DAOProduct{}, nil
	}
	return result[0], nil
}

func (r sqlrepository) Find(ctx context.Context, filter ProductFilterRequest) (products []DAOProduct, err error) {
	log.Println("Filters: Id: " + strconv.FormatInt(filter.ProductId, 10) + " Name: " + filter.ProductName)
	query := "SELECT ProductId,Name,Description,Price,SKU FROM Products WHERE 1=1"

	if filter.ProductId > 0 {
		query += " AND ProductId = " + strconv.FormatInt(filter.ProductId, 10)
	} else {
		if filter.ProductName != "" {
			query += " AND Name Like '%" + filter.ProductName + "%'"
		}
	}

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Could not fetch data from the database with error: ", err)
		return nil, errors.New("Something get wrong with data base")
	}

	return scanProductRows(rows)
}

func (r sqlrepository) GetAll(ctx context.Context) (products []DAOProduct, err error) {
	query := "SELECT ProductId,Name,Description,Price,SKU FROM Products"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Could not fetch data from the database with error: ", err)
		return nil, errors.New("Something get wrong with data base")
	}

	return scanProductRows(rows)
}

func (r sqlrepository) Update(ctx context.Context, product DAOProduct) error {

	row := r.db.QueryRowContext(ctx, "UPDATE Products SET Name = ?, Description = ?, Price = ?, SKU = ?, UpdatedOn = ? WHERE ProductId = ?",
		*&product.Name, *&product.Description, *&product.Price, *&product.SKU, time.Now())

	if row.Err() != nil {
		log.Fatal("Could not update records with error: ", row.Err())
	}
	return row.Err()
}

func (r sqlrepository) Insert(ctx context.Context, product DAOProduct) error {
	query := "INSERT INTO Products(Name, Description, Price, SKU, CreatedOn) VALUES (?, ?, ?, ?, ?)"
	row := r.db.QueryRowContext(ctx, query,
		product.Name, product.Description, product.Price, product.SKU, time.Now())

	if row.Err() != nil {
		log.Fatal("Could not fetch data from the database with error: ", row.Err())
	}
	return row.Err()
}

func scanProductRows(rows *sql.Rows) (result []DAOProduct, err error) {
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatal("Something went wrong when scanning Rows with error: ", err)
		}
	}(rows)

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var actual DAOProduct
		if err = rows.Scan(&actual.ProductId, &actual.Name, &actual.Description, &actual.Price,
			&actual.SKU); err != nil {
			log.Fatal("Something went wrong when scanning Rows with error: ", err)
			return nil, err
		}
		result = append(result, actual)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Something whent wrong when scanning Rows with error: ", err)
		return nil, err
	}
	return result, nil
}

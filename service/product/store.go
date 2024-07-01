package product

import (
	"database/sql"
	"go-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (store *Store) CreateProduct(product types.Product) error {
	_, err := store.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		product.Name, product.Description, product.Image, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) GetProducts() ([]types.Product, error) {
	rows, err := store.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		product, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

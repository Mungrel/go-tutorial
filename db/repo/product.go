package repo

import (
	"database/sql"
	"meme/db"
	"meme/types"

	"github.com/gofrs/uuid"
)

// CreateProduct creates a product in the DB.
// It will generate the ID itself.
func CreateProduct(product types.Product) (types.Product, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return types.Product{}, err
	}

	product.ID = id.String()

	const insertProduct = `
		INSERT INTO product (
			id,
			name,
			price
		) VALUES (
			:id,
			:name,
			:price
		)`

	_, err = db.Client().NamedExec(insertProduct, product)
	if err != nil {
		return types.Product{}, err
	}

	return product, nil
}

// GetProductByID returns the product associated with the provided ID.
// It will return nil (and no error) if no product was found.
func GetProductByID(id string) (*types.Product, error) {
	const getProduct = `
		SELECT id, name, price
		FROM product
		WHERE id = ?`

	var product types.Product
	err := db.Client().Get(&product, getProduct, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, err
}

// UpdateProduct updates the product associated with the provided ID with the provided product value.
// It will return nil if there was no product for that ID.
func UpdateProduct(id string, product types.Product) (*types.Product, error) {
	const updateProduct = `
		UPDATE product SET
			name = :name,
			price = :price
		WHERE id = :id`

	product.ID = id

	result, err := db.Client().NamedExec(updateProduct, product)
	if err != nil {
		return nil, err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAff == 0 {
		return nil, nil
	}

	return &product, nil
}

// DeleteProduct deletes the product associated with the provided ID.
// It will *not* return an error if no product was found for that ID.
func DeleteProduct(id string) error {
	const delete = `
		DELETE
		FROM product
		WHERE id = ?`

	_, err := db.Client().Exec(delete, id)
	return err
}

// ListProducts lists the first 10 products in the DB.
func ListProducts() ([]types.Product, error) {
	const list = `
		SELECT id, name, price
		FROM product
		LIMIT 10`

	products := []types.Product{}
	err := db.Client().Select(&products, list)
	return products, err
}

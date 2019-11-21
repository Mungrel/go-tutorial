package handle

import (
	"meme/db/repo"
	"meme/types"
	"net/http"

	"github.com/bouk/httprouter"
)

// GetProduct handles product GET requests.
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	product, err := repo.GetProductByID(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, product, http.StatusOK)
}

// CreateProduct handles POST requests.
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product types.Product
	err := readBody(r, &product)
	if err != nil {
		respondWithError(w, err)
		return
	}

	savedProduct, err := repo.CreateProduct(r.Context(), product)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, savedProduct, http.StatusCreated)
}

// UpdateProduct handles product PUT requests.
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product types.Product
	err := readBody(r, &product)
	if err != nil {
		respondWithError(w, err)
		return
	}

	id := httprouter.GetParam(r, "id")

	updatedProduct, err := repo.UpdateProduct(r.Context(), id, product)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, updatedProduct, http.StatusOK)
}

// DeleteProduct handles product DELETE requests.
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := httprouter.GetParam(r, "id")

	err := repo.DeleteProduct(r.Context(), id)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithStatus(w, http.StatusOK)
}

// ListProducts handles product list GET requests.
func ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := repo.ListProducts(r.Context())
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, products, http.StatusOK)
}

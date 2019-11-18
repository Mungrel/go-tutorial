package types

import "github.com/shopspring/decimal"

// Product represents a product in the DB and API.
type Product struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
}

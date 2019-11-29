package repo

import (
	"context"
	"meme/db"
	"meme/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe("Product Repo", func() {
	var ctx context.Context
	BeforeEach(func() {
		db.TablesDir = "../tables/"
		ctx = db.ContextWithDB(context.Background(), db.Client())
	})

	Describe("#CreateProduct", func() {
		It("should create a product without error", func() {
			p := types.Product{
				Name:  "Hello",
				Price: decimal.NewFromFloat(1.50),
			}
			savedProduct, err := CreateProduct(ctx, p)
			Expect(err).ToNot(HaveOccurred())
			Expect(savedProduct.ID).ToNot(BeEmpty())

			p.ID = savedProduct.ID
			Expect(savedProduct).To(Equal(p))
		})
	})
})

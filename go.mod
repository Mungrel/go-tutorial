module meme

go 1.13

require (
	github.com/bouk/httprouter v0.0.0-20160817010721-ee8b3818a7f5
	github.com/go-sql-driver/mysql v1.4.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/jmoiron/sqlx v1.2.0
	github.com/onsi/ginkgo v1.10.3
	github.com/onsi/gomega v1.7.1

	// Pinned because of https://github.com/golang/go/issues/34315
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/stretchr/testify v1.4.0
	golang.org/x/sys v0.0.0-20191128015809-6d18c012aee9 // indirect
)

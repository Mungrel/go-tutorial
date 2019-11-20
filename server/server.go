package server

import (
	"meme/db"
	"meme/handle"
	"net/http"

	"github.com/bouk/httprouter"
)

func New() http.Handler {
	router := httprouter.New()

	router.GET("/api/product/:id", handle.GetProduct)
	router.POST("/api/product", handle.CreateProduct)
	router.PUT("/api/product/:id", handle.UpdateProduct)
	router.DELETE("/api/product/:id", handle.DeleteProduct)
	router.GET("/api/products", handle.ListProducts)

	return &Server{
		router: router,
	}
}

type Server struct {
	router http.Handler
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = db.ContextWithDB(ctx, db.Client())
	server.router.ServeHTTP(w, r.WithContext(ctx))
}

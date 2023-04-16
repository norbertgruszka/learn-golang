package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/norbertgruszka/learn-golang/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *chi.Mux) {
	router.Post("/book/", controllers.CreateBook)
	router.Get("/book/", controllers.GetBook)
	router.Get("/book/{bookId}", controllers.GetBookById)
	router.Put("/book/{bookId}", controllers.UpdateBook)
	router.Delete("/book/{bookId}", controllers.DeleteBook)
}

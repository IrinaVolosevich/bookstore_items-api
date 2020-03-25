package application

import (
	"bookstore_items-api/src/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)

	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
}

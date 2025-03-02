package controllers

import (
	"encoding/json"
	"go-bookstore/pkg/models"
	"net/http"

	"github.com/nishanthanr151998/go-bookstore/pkg/models"
)

var NewBook models.Book

func getBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

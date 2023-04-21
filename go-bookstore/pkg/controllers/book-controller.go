package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/norbertgruszka/learn-golang/go-bookstore/pkg/models"
	"github.com/norbertgruszka/learn-golang/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	if id := chi.URLParam(r, "bookId"); id != "" {
		ID, err := strconv.ParseInt(id, 0, 0)
		if err != nil {
			fmt.Println("error while parsing")
		}
		bookDetails, _ := models.GetBookById(ID)
		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletBook(w http.ResponseWriter, r *http.Request) {
	if id := chi.URLParam(r, "bookId"); id != "" {
		ID, err := strconv.ParseInt(id, 0, 0)
		if err != nil {
			fmt.Println("error while parsing")
		}
		book := models.DeleteBook(ID)
		res, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	if id := chi.URLParam(r, "bookId"); id != "" {
		ID, err := strconv.ParseInt(id, 0, 0)
		if err != nil {
			fmt.Println("error while parsing")
		}

		bookDetails, db := models.GetBookById(ID)
	}
}

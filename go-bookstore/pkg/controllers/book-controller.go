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

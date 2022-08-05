package handler

import "github.com/saidakhmatov/catalog_of_books/storage"

type handler struct {
	strg storage.StorageI
}

func NewHandler(strg storage.StorageI) *handler {
	return &handler{
		strg: strg,
	}
}

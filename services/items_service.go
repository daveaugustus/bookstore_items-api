package services

import (
	"net/http"

	"github.com/davetweetlive/bookstore_items-api/domain/items"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

var (
	ItemService itemServiceInterface = &itemsServie{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsServie struct{}

func (s *itemsServie) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not implemented", nil)
}

func (s *itemsServie) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not implemented", nil)
}

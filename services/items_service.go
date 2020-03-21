package services

import (
	"bookstore_items-api/domain/items"
	"github.com/IrinaVolosevich/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, nil

	itemRequest.Save()

	return nil, nil
}

func (s *itemsService) Get(itemId string) (*items.Item, *rest_errors.RestErr) {
	return nil, nil
}

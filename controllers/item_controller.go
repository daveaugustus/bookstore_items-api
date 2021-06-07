package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/davetweetlive/bookstore_items-api/domain/items"
	"github.com/davetweetlive/bookstore_items-api/services"
	"github.com/davetweetlive/bookstore_oauth-go/oauth"
)

var (
	ItemController itemControllerInterface = &itemController{}
)

type itemControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemController struct{}

func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error json to the user
		return
	}

	req := oauth.GetCallerId(r)
	st := strconv.FormatInt(req, 10)
	item := items.Item{
		Seller: st,
	}

	result, err := services.ItemService.Create(item)
	if err != nil {
		// TODO: Return error json to the user
		return
	}
	fmt.Println(result)
	// TODO: Return created item as JSON with HTTP status 201 - Created.
}

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {

}

package splio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ErrBadRequest = errors.New("ErrBadRequest")
var ErrForbidden = errors.New("ErrForbidden")
var ErrNotFound = errors.New("ErrNotFound")

const apiUrl = "https://api.splio.com"

type Client interface {
	Authenticate() *ApiError

	CreateContact(contact *Contact) *ApiError
	DeleteContact(contactKey string) *ApiError
	EditContact(contact *Contact) *ApiError
	GetContact(contactKey string) (*Contact, *ApiError)
	ListContact(perPage, pageNumber int, searchFields []SearchField) (*SearchResult, *ApiError)
	OptOutContact(contactKeys []string) []*ApiError
	SubscribeContact(contactKey string, listIds []int) *ApiError
	UnSubscribeContact(contactKey string, listIds []int) *ApiError

	CreateOrder(order *Order) *ApiError
	EditOrder(order *Order) *ApiError
	GetOrder(orderKey string) (*Order, *ApiError)

	CreateProduct(product *Product) *ApiError
	EditProduct(product *Product) *ApiError
	GetProduct(productKey string) (*Product, *ApiError)
}

type client struct {
	universe string
	key      string
	token    string
}

func NewClient(universe, key string) Client {
	return &client{
		universe: universe,
		key:      key,
	}
}

func (c *client) Authenticate() *ApiError {
	var result AuthenticateResponse

	apiError := c.call(http.MethodPost, apiUrl+"/authenticate", &Authenticate{ApiKey: c.key}, &result, false)
	if apiError != nil {
		return apiError
	}

	c.token = "Bearer" + result.Token

	return nil
}

func (c *client) CreateContact(contact *Contact) *ApiError {
	var result Contact

	apiError := c.call(http.MethodPost, apiUrl+"/data/contacts", contact, &result, true)
	if apiError != nil {
		return apiError
	}

	contact = &result

	return nil
}

func (c *client) DeleteContact(contactKey string) *ApiError {
	uri := fmt.Sprintf(apiUrl+"/data/contacts/%s", contactKey)

	apiError := c.call(http.MethodDelete, uri, &struct{}{}, nil, true)
	if apiError != nil {
		return apiError
	}
	return nil
}

func (c *client) EditContact(contact *Contact) *ApiError {
	var result Contact
	uri := fmt.Sprintf(apiUrl+"/data/contacts/%s", contact.Email)

	contact.Id = nil

	apiError := c.call(http.MethodPatch, uri, contact, &result, true)
	if apiError != nil {
		return apiError
	}

	contact = &result

	return nil
}

func (c *client) GetContact(contactKey string) (*Contact, *ApiError) {
	var result Contact
	uri := fmt.Sprintf(apiUrl+"/data/contacts/%s", contactKey)

	apiError := c.call(http.MethodGet, uri, &struct{}{}, &result, true)
	if apiError != nil {
		return nil, apiError
	}

	return &result, nil
}

func (c *client) ListContact(perPage, pageNumber int, searchFields []SearchField) (*SearchResult, *ApiError) {
	var result SearchResult

	apiError := c.call(http.MethodGet, apiUrl+"/data/contacts", &SearchBody{PerPage: perPage, PageNumber: pageNumber, Fields: searchFields}, &result, true)
	if apiError != nil {
		return nil, apiError
	}

	return &result, nil
}

func (c *client) OptOutContact(contactKeys []string) []*ApiError {
	var result MultiStatus

	body := struct {
		contacts []string
	}{contacts: contactKeys}

	apiError := c.call(http.MethodPost, apiUrl+"/data/contacts/optout", &body, &result, true)
	if apiError != nil {
		return []*ApiError{apiError}
	}

	if result.ErrorCount > 0 {
		var response []*ApiError
		for _, item := range result.Items {
			var err error
			if item.Code == http.StatusBadRequest {
				err = ErrBadRequest
			} else if item.Code == http.StatusForbidden {
				err = ErrForbidden
			} else if item.Code == http.StatusNotFound {
				err = ErrNotFound
			}

			if err != nil {
				response = append(response, &ApiError{Err: err, StatusCode: item.Code, Errors: []ApiErrorDesc{{ErrorDescription: item.Description}}})
			}
		}
		return response
	}

	return nil
}

func (c *client) SubscribeContact(contactKey string, listIds []int) *ApiError {
	apiError := c.subscribeHandling("/data/contacts/%s/lists/subscribe", contactKey, listIds)
	if apiError != nil {
		return apiError
	}

	return nil
}

func (c *client) UnSubscribeContact(contactKey string, listIds []int) *ApiError {
	apiError := c.subscribeHandling("/data/contacts/%s/lists/unsubscribe", contactKey, listIds)
	if apiError != nil {
		return apiError
	}

	return nil
}

func (c *client) subscribeHandling(uri, contactKey string, listIds []int) *ApiError {
	uri = fmt.Sprintf(apiUrl+uri, contactKey)
	var result Contact

	body := struct {
		ListIds []int `json:"list_ids"`
	}{ListIds: listIds}

	apiError := c.call(http.MethodPost, uri, &body, &result, true)
	if apiError != nil {
		return apiError
	}

	return nil
}

func (c *client) CreateOrder(order *Order) *ApiError {
	var result Order

	apiError := c.call(http.MethodPost, apiUrl+"/data/v1/orders", order, &result, true)
	if apiError != nil {
		return apiError
	}

	order = &result

	return nil
}

func (c *client) EditOrder(order *Order) *ApiError {
	var result Order
	uri := fmt.Sprintf(apiUrl+"/data/v2/orders/%s", *order.ExternalId)

	order.ExternalId = nil

	apiError := c.call(http.MethodPatch, uri, order, &result, true)
	if apiError != nil {
		return apiError
	}

	order = &result

	return nil
}

func (c *client) GetOrder(orderKey string) (*Order, *ApiError) {
	var result Order
	uri := fmt.Sprintf(apiUrl+"/data/v1/orders/%s", orderKey)

	apiError := c.call(http.MethodGet, uri, &struct{}{}, &result, true)
	if apiError != nil {
		return nil, apiError
	}

	return &result, nil
}

func (c *client) CreateProduct(product *Product) *ApiError {
	var result Product

	apiError := c.call(http.MethodPost, apiUrl+"/data/v1/products", product, &result, true)
	if apiError != nil {
		return apiError
	}

	product = &result

	return nil
}

func (c *client) EditProduct(product *Product) *ApiError {
	var result Product
	uri := fmt.Sprintf(apiUrl+"/data/v1/products/%s", *product.ExternalId)

	product.ExternalId = nil

	apiError := c.call(http.MethodPatch, uri, product, &result, true)
	if apiError != nil {
		return apiError
	}

	product = &result

	return nil
}

func (c *client) GetProduct(productKey string) (*Product, *ApiError) {
	var result Product
	uri := fmt.Sprintf(apiUrl+"/data/v1/products/%s", productKey)

	apiError := c.call(http.MethodGet, uri, &struct{}{}, &result, true)
	if apiError != nil {
		return nil, apiError
	}

	return &result, nil
}

func (c *client) call(method, url string, body, result interface{}, needAuthentication bool) *ApiError {
	// Json encode body
	if body == nil {
		body = ""
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return &ApiError{Err: err}
	}

	httpClient := http.DefaultClient
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		return &ApiError{Err: err}
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	if needAuthentication {
		if c.token == "" {
			apiError := c.Authenticate()
			if apiError != nil {
				return apiError
			}
		}

		if c.token != "" {
			req.Header.Set("authorization", c.token)
		}
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return &ApiError{Err: err}
	}

	if response.StatusCode == http.StatusUnauthorized {
		// Clear token and retry call
		c.token = ""
		return c.call(method, url, body, result, needAuthentication)
	} else {
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return &ApiError{Err: err}
		}

		defer func() {
			err := response.Body.Close()
			if err != nil {
				panic(fmt.Sprintf("can not close response body: %v\n", err))
			}
		}()

		err = nil
		if response.StatusCode == http.StatusBadRequest {
			err = ErrBadRequest
		} else if response.StatusCode == http.StatusForbidden {
			err = ErrForbidden
		} else if response.StatusCode == http.StatusNotFound {
			err = ErrNotFound
		}

		if err == nil {
			if result != nil {
				err = json.Unmarshal(responseBody, result)
				if err != nil {
					return &ApiError{Err: err}
				}
			}
		} else {
			var apiError ApiError
			err2 := json.Unmarshal(responseBody, &apiError)
			if err2 != nil {
				return &ApiError{Err: err2}
			}
			apiError.Err = err
			return &apiError
		}
	}

	return nil
}

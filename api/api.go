package api

import (
	"fmt"
	"freshdesk/models"
)

//FreshDeskClient stores the api key
type FreshDeskClient struct {
	API
}

//NewClient does bla bla
func NewClient(domain, username, password string, secure bool) FreshDeskClient {
	protocol := "http"
	if secure {
		protocol = "https"
	}
	api := NewAPI(protocol, domain, username, password)
	return FreshDeskClient{API: api}
}

// UserCreate does this
func (client *FreshDeskClient) UserCreate(name, email string) (models.User, error) {
	var userResponse models.UserResponse
	userResponse.User = models.User{Name: name, Email: email}
	requestUrl := client.BaseUrl() + fmt.Sprintf("/contacts.json")
	err := client.DoWithResultEx(requestUrl, POST, userResponse.Converts(), &userResponse, connectTimeOut, readWriteTimeout, CONTENT_TYPE_APPLICATION_JSON)
	if err != nil {
		return models.User{}, err
	}

	return userResponse.User, err
}

//UserView does
func (client *FreshDeskClient) UserView(id int) (models.User, error) {
	var userResponse models.UserResponse = models.UserResponse{}
	requestUrl := client.BaseUrl() + fmt.Sprintf("/contacts/%v.json", id)
	err := client.DoWithResult(requestUrl, GET, &userResponse)
	if err != nil {
		return models.User{}, err
	}
	return userResponse.User, err
}

func (client *FreshDeskClient) UserDelete(id int) (bool, error) {
	requestUrl := client.BaseUrl() + fmt.Sprintf("/contacts/%v.json", id)
	err := client.DoWithResult(requestUrl, DELETE, nil)
	return err == nil, err
}

func (client *FreshDeskClient) CustomerCreate(name, domains, description string) (Customer, error) {
	var customerResponse CustomerResponse = CustomerResponse{}
	customerResponse.Customer = Customer{Name: name, Domains: domains, Description: description}
	requestUrl := client.BaseUrl() + fmt.Sprintf("/customer.json")
	err := client.DoWithResultEx(requestUrl, POST, customerResponse.Json(), &customerResponse, connectTimeOut, readWriteTimeout, CONTENT_TYPE_APPLICATION_JSON)
	if err != nil {
		return Customer{}, err
	}
	return customerResponse.Customer, err
}

func (client *FreshDeskClient) CustomerList(filter string) ([]Customer, error) {
	var customerResponses []CustomerResponse
	requestUrl := client.BaseUrl() + fmt.Sprintf("/customers.json")
	if filter != "" {
		requestUrl = requestUrl + fmt.Sprintf("?letter=%s", filter)
	}
	err := client.DoWithResult(requestUrl, GET, &customerResponses)
	var customers []Customer
	if err != nil {
		return customers, err
	}
	for _, response := range customerResponses {
		customers = append(customers, response.Customer)
	}
	return customers, err

}

func (client *FreshDeskClient) CustomerView(id int) (Customer, error) {
	var customerResponses CustomerResponse
	requestUrl := client.BaseUrl() + fmt.Sprintf("/customers/%v.json", id)
	err := client.DoWithResult(requestUrl, GET, &customerResponses)
	if err != nil {
		return Customer{}, err
	}
	return customerResponses.Customer, err
}

func (client *FreshDeskClient) CustomerDelete(id int) (bool, error) {
	requestUrl := client.BaseUrl() + fmt.Sprintf("/customers/%v.json", id)
	err := client.DoWithResult(requestUrl, DELETE, nil)
	return err == nil, err
}

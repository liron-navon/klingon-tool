package stapi

import (
	"fmt"
	"net/http"
)

// ApiUrl - the base url for stapi rest api
const ApiUrl = "http://stapi.co/api/v1/rest"

// Client - the stapi app
type Client struct {
	ApiUrl     string
	HttpClient *http.Client
	Character  Entity
}

// New - create a new stapi client
func New(httpClient *http.Client) Client {
	c := Client{
		ApiUrl:     ApiUrl,
		HttpClient: httpClient,
	}

	c.Character = Entity{
		ApiUrl: fmt.Sprintf("%s/character", c.ApiUrl),
		Client: c.HttpClient,
	}

	return c
}

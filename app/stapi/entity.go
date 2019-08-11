package stapi

import (
	"encoding/json"
	"fmt"
	"klingon-tool/app/stapi/types"
	"net/http"
	"net/url"
	"strings"
)

const (
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)

type Entity struct {
	Client *http.Client
	ApiUrl string
}

// Search - calls a POST request for 'search', for a stapi resource
func (e *Entity) Search(query url.Values) (response types.SearchResponse, httpResponse *http.Response, err error) {
	apiURL := fmt.Sprintf("%s/search", e.ApiUrl)
	searchResponse := types.SearchResponse{}

	resp, err := e.Client.Post(
		apiURL,
		ContentTypeFormURLEncoded,
		strings.NewReader(query.Encode()),
	)
	if err != nil {
		return searchResponse, resp, err
	}

	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return searchResponse, resp, err
	}

	return searchResponse, resp, nil
}

// Fetch - calls a GET request for a stapi resource (requires the 'uid' query parameter)
func (e *Entity) Fetch(query url.Values) (response types.FetchResponse, httpResponse *http.Response, err error) {
	apiURL := fmt.Sprintf("%s?%s", e.ApiUrl, query.Encode())
	fetchResponse := types.FetchResponse{}

	resp, err := e.Client.Get(apiURL)

	if err != nil {
		return fetchResponse, resp, err
	}

	err = json.NewDecoder(resp.Body).Decode(&fetchResponse)
	if err != nil {
		return fetchResponse, resp, err
	}

	return fetchResponse, resp, nil
}

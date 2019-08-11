package stapi

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// createMockHttpClient - creates a simple mock server that returns a string value
func createMockHttpClient(t *testing.T, returnValue string) (*http.Client, *httptest.Server) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(returnValue))
		if err != nil {
			t.Fatal(err)
		}
	})
	server := httptest.NewServer(h)

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, server.Listener.Addr().String())
			},
		},
	}

	return client, server
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name           string
		collectionName string
		expectingError bool
	}{
		{"Worf", "characters", false},
		{"Klingon", "species", false},
	}

	for _, test := range tests {
		client, server := createMockHttpClient(t, fmt.Sprintf(`{
			"%s": [{ "uid": "X", "name": "%s" }]
		}`, test.collectionName, test.name))
		stapi := New(client)

		query := url.Values{}
		query.Add("name", test.name)

		resp, err := stapi.Character.Search(query)
		if err != nil {
			if !test.expectingError {
				t.Fatal(err)
			}
			t.Fatal(err)
		}
		switch test.collectionName {
		case "species":
			require.Equal(t, resp.Species[0].Name, test.name)
		case "characters":
			require.Equal(t, resp.Characters[0].Name, test.name)
		}

		server.Close()
	}
}

func TestFetch(t *testing.T) {
	tests := []struct {
		name           string
		uid            string
		expectingError bool
	}{
		{"Worf", "CHMA0000009023", false},
		{"Nyota Uhura", "CHMA0000115364", false},
	}

	for _, test := range tests {
		client, server := createMockHttpClient(t, "{}")
		stapi := New(client)

		query := url.Values{}
		query.Add("uid", test.uid)

		_, err := stapi.Character.Fetch(query)
		if err != nil {
			if !test.expectingError {
				t.Fatal(err)
			}
			t.Fatal(err)
		}
		server.Close()
	}
}

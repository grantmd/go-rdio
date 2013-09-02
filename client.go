package rdio

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	RDIO_API_ENDPOINT   = "https://api.rdio.com/1"
	RDIO_OAUTH_ENDPOINT = "https://api.rdio.com/oauth"
)

type Client struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
}

func (c *Client) Call(method string, params map[string][]string) []byte {
	params["method"] = []string{method}
	// TODO: JSON decode
	return c.SignedPost(RDIO_API_ENDPOINT, params)
}

func (c *Client) SignedPost(url string, params map[string][]string) []byte {
	// Make call
	resp, err := http.PostForm(url, params)
	if err != nil {
		log.Fatal(err)
	}

	// Read body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Return
	return body
}

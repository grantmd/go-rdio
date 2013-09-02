package rdio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	RDIO_API_ENDPOINT   = "http://api.rdio.com/1/"
	RDIO_OAUTH_ENDPOINT = "http://api.rdio.com/oauth"
)

type Client struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
}

func (c *Client) Call(method string, params map[string][]string) interface{} {
	params["method"] = []string{method}
	body, err := c.SignedPost(RDIO_API_ENDPOINT, params)
	if err != nil {
		log.Fatal(err)
	}

	// parse into json
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func (c *Client) SignedPost(url string, params map[string][]string) ([]byte, error) {
	// Make call
	resp, err := http.PostForm(url, params)
	if err != nil {
		return nil, err
	}

	// Make sure we close the body stream no matter what
	defer resp.Body.Close()

	// Check status code
	switch resp.StatusCode {
	default:
		return nil, fmt.Errorf("Unknown status code: %d", resp.StatusCode)
	case 403:
		return nil, errors.New("Developer Inactive")
	case 200:

	}

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	str := string(body)
	fmt.Println(str)

	// Return
	return body, nil
}

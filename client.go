package rdio

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	RDIO_API_ENDPOINT   = "http://api.rdio.com/1/"
	RDIO_OAUTH_ENDPOINT = "http://api.rdio.com/oauth"
)

type Client struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

func (c *Client) Call(method string, params url.Values) (interface{}, error) {
	params["method"] = []string{method}
	body, err := c.SignedPost(RDIO_API_ENDPOINT, params)
	if err != nil {
		return nil, err
	}

	// parse into json
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (c *Client) SignedPost(postUrl string, params url.Values) ([]byte, error) {
	// Sign the params
	auth := c.Sign(postUrl, params)
	fmt.Println(auth)

	// Make call
	resp, err := http.PostForm(postUrl, params)
	if err != nil {
		return nil, err
	}

	// Make sure we close the body stream no matter what
	defer resp.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check status code
	switch resp.StatusCode {
	default:
		str := string(body)
		return nil, fmt.Errorf("Unknown status code: %d, %s", resp.StatusCode, str)
	case 400:
		return nil, errors.New("Bad Request")
	case 403:
		return nil, errors.New("Developer Inactive")
	case 200:

	}

	// Return
	return body, nil
}

func (c *Client) Sign(signUrl string, params url.Values) string {
	params["oauth_version"] = []string{"1.0"}
	params["oauth_timestamp"] = []string{strconv.FormatInt(time.Now().Unix(), 10)}
	params["oauth_nonce"] = []string{strconv.FormatInt(rand.Int63(), 10)}
	params["oauth_signature_method"] = []string{"HMAC-SHA1"}
	params["oauth_consumer_key"] = []string{c.ConsumerKey}

	// The consumer secret is the first half of the HMAC-SHA1 key
	hmacKey := c.ConsumerSecret + "&"

	if c.Token != "" {
		// Include a token in params
		params["oauth_token"] = []string{c.Token}
		// and the token secret in the HMAC-SHA1 key
		hmacKey += c.TokenSecret
	}

	// sort the params by key
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sorted := url.Values{}
	for _, k := range keys {
		sorted.Add(k, params.Get(k))
	}

	// Build the signature base string
	signatureBaseString := []byte("POST&" + url.QueryEscape(signUrl) + "&" + sorted.Encode())

	// Calculate HMAC-SHA1
	mac := hmac.New(sha1.New, []byte(hmacKey))
	mac.Write(signatureBaseString)
	oauthSignature := base64.URLEncoding.EncodeToString(mac.Sum(nil))

	// Build the Authorization header
	authorizationParams := url.Values{}
	authorizationParams.Add("oauth_signature", oauthSignature)

	// List of params that must be included in the header, if present
	for _, k := range keys {
		switch k {
		case "oauth_version",
			"oauth_timestamp",
			"oauth_nonce",
			"oauth_signature_method",
			"oauth_signature",
			"oauth_consumer_key",
			"oauth_token":

			authorizationParams.Add(k, params.Get(k))
		}
	}

	return "OAuth " + strings.Replace(authorizationParams.Encode(), "&", ", ", -1)
}

func (c *Client) StartAuth() ([]byte, error) {
	params := url.Values{
		"oauth_callback": []string{"oob"},
	}

	body, err := c.SignedPost("http://api.rdio.com/oauth/request_token", params)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) CompleteAuth(verifier string) ([]byte, error) {
	params := url.Values{
		"oauth_verifier": []string{"verifier"},
	}

	body, err := c.SignedPost("http://api.rdio.com/oauth/access_token", params)
	if err != nil {
		return nil, err
	}

	return body, nil
}

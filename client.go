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
	RDIO_OAUTH_ENDPOINT = "http://api.rdio.com/oauth/"
)

type Client struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
	httpClient     *http.Client
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

	// Build HTTP client
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	postBody := params.Encode()
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(postBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	// Sign the params
	auth := c.Sign(postUrl, params)
	//fmt.Println(auth)

	req.Header.Set("Authorization", auth)
	resp, err := c.httpClient.Do(req)
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
		return nil, fmt.Errorf("Unknown status code: %d", resp.StatusCode)
	case 400:
		return nil, errors.New("Bad Request")
	case 401:
		return nil, errors.New("Invalid Signature")
	case 403:
		return nil, errors.New("Developer Inactive")
	case 200:

	}

	// Return
	return body, nil
}

func (c *Client) Sign(signUrl string, params url.Values) string {
	rand.Seed(time.Now().UnixNano())
	params["oauth_version"] = []string{"1.0"}
	params["oauth_timestamp"] = []string{strconv.FormatInt(time.Now().Unix(), 10)}
	params["oauth_nonce"] = []string{strconv.FormatInt(rand.Int63n(1000000), 10)}
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
	signatureBaseString := []byte("POST&" + url.QueryEscape(signUrl) + "&" + url.QueryEscape(sorted.Encode()))

	// Calculate HMAC-SHA1
	mac := hmac.New(sha1.New, []byte(hmacKey))
	mac.Write(signatureBaseString)
	oauthSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	//fmt.Println(oauthSignature)

	// Build the Authorization header
	authorizationParams := url.Values{}
	authorizationParams.Add("oauth_signature", `"`+oauthSignature+`"`)

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

			authorizationParams.Add(k, `"`+params.Get(k)+`"`)
		}
	}

	return "OAuth " + strings.Replace(strings.Replace(authorizationParams.Encode(), "&", ", ", -1), "%22", `"`, -1)
}

func (c *Client) StartAuth() (url.Values, error) {
	// Request token
	params := url.Values{
		"oauth_callback": []string{"oob"},
	}

	body, err := c.SignedPost(RDIO_OAUTH_ENDPOINT+"request_token", params)
	if err != nil {
		return nil, err
	}

	// Parse response to extract login url, request token, and request secret
	m, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) CompleteAuth(verifier string) (url.Values, error) {
	// Request exchange for access token
	params := url.Values{
		"oauth_verifier": []string{"verifier"},
	}

	body, err := c.SignedPost(RDIO_OAUTH_ENDPOINT+"access_token", params)
	if err != nil {
		return nil, err
	}

	// Parse response to extract access token and secret
	m, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	return m, nil
}

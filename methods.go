package rdio

import (
	"encoding/json"
	"errors"
)

func (c *Client) GetNewReleases() ([]Album, error) {
	params := make(map[string][]string)
	body, err := c.Call("getNewReleases", params)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Status string
		Result []Album
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return nil, errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
}

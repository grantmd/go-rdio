//
// This is a list of methods that map to Rdio API methods. Calling them returns a proper
// Go struct of the data, which is probably preferable to getting a blob of JSON or
// a byte stream.
//
// Unfortunately, there is a bit of repetition in this code (the Response struct and
// checking the Status == "ok" response), but I don't know how to consolidate it
//

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

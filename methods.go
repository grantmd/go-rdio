package rdio

import (
	"fmt"
	"reflect"
)

func (c *Client) GetNewReleases() ([]Album, error) {
	params := make(map[string][]string)
	data, err := c.Call("getNewReleases", params)
	if err != nil {
		return nil, err
	}

	var albums []Album
	err = parseResult(data, &albums)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

// Private function to take the API response JSON and stuff it into whatever struct we want
func parseResult(data map[string]interface{}, v interface{}) error {
	fmt.Println(reflect.ValueOf(v).Kind())
	result := data["result"].([]interface{})
	for _, item := range result {
		fmt.Println(item)
	}
	return nil
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationInfo(location string) (LocationInfo, error) {
	url := baseURL + "/location-area/" + location

	locationInfo := LocationInfo{}

	val, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &locationInfo)
		if err != nil {
			return LocationInfo{}, err
		}
		return locationInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationInfo{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locationInfo)
	if err != nil {
		return LocationInfo{}, err
	}

	return locationInfo, nil
}

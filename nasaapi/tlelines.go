package nasaapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NASATLEAPIResponse struct {
	Member []TLE `json:"member"`
}

type TLE struct {
	ID          string `json:"@id"`
	Type        string `json:"@type"`
	SatelliteID int    `json:"satelliteId"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
}

func CallGeneralNASATLEInformation() ([]TLE, error) {
	baseURL := "https://tle.ivanstanojevic.me/api/tle/"

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Error sending get req to NASA TLE api: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading resp body from NASA TLE api: %v", err)
	}

	var tles NASATLEAPIResponse
	err = json.Unmarshal(body, &tles)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling json reply from NASA TLE api: %v", err)
	}

	return tles.Member, nil
}
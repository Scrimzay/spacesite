package spacexapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpaceXCrewAPI struct {
	Name string `json:"name"`
	Agency string `json:"agency"`
	Status string `json:"status"`
	ID string `json:"id"`
}

func CallGeneralSpaceXCrewInformation() ([]SpaceXCrewAPI, error) {
	baseURL := "https://api.spacexdata.com/v4/crew"

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Could not send request to base url for spacex crew: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read body resp from spacex crew: %v", err)
	}

	var crewInfo []SpaceXCrewAPI
	err = json.Unmarshal(body, &crewInfo)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal json from spacex crew: %v", err)
	}

	return crewInfo, nil
}
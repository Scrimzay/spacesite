package spacexapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpaceXCompanyAPI struct {
	Name string `json:"name"`
	Founder string `json:"founder"`
	Founded int `json:"founded"`
	Employees int `json:"employees"`
	Vehicles int `json:"vehicles"`
	LaunchSites int `json:"launch_sites"`
	TestSites int `json:"test_sites"`
	CEO string `json:"ceo"`
	CTO string `json:"cto"`
	COO string `json:"coo"`
	CTOPropulsion string `json:"cto_propulsion"`
	Valuation int `json:"valuation"`
	Summary string `json:"summary"`
	ID string `json:"id"`
}

func CallGeneralSpaceXCompanyInformation() (*SpaceXCompanyAPI, error) {
	baseURL := "https://api.spacexdata.com/v4/company"

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Could not send request to base url for spacex company: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read body resp from spacex company: %v", err)
	}

	var companyInfo SpaceXCompanyAPI
	err = json.Unmarshal(body, &companyInfo)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal json from spacex company: %v", err)
	}

	return &companyInfo, nil
}
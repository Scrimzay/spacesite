package spacexapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpaceXRoadsterAPI struct {
	Name string `json:"name"`
	LaunchDateUTC string `json:"launch_date_utc"`
	LaunchMassKG int `json:"launch_mass_kg"`
	LaunchMassLBS int `json:"launch_mass_lbs"`
	OrbitType string `json:"orbit_type"`
	Longitude float32 `json:"longitude"`
	PeriodDays float32 `json:"period_days"`
	SpeedKPH float32 `json:"speed_kph"`
	SpeedMPH float32 `json:"speed_mph"`
	EarthDistanceKM float32 `json:"earth_distance_km"`
	EarthDistanceMI float32 `json:"earth_distance_mi"`
	MarsDistanceKM float32 `json:"mars_distance_km"`
	MarsDistanceMI float32 `json:"mars_distance_mi"`
	Details string `json:"details"`
	ID string `json:"id"`
}

func CallGeneralSpaceXRoadsterInformation() (*SpaceXRoadsterAPI, error) {
	baseURL := "https://api.spacexdata.com/v4/roadster"
    resp, err := http.Get(baseURL)
    if err != nil {
        return nil, fmt.Errorf("could not send GET to SpaceX roadster API: %v", err)
    }
    defer resp.Body.Close()

    // Read response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response for SpaceX roadster: %v", err)
    }

    var roadster SpaceXRoadsterAPI
    err = json.Unmarshal(body, &roadster)
    if err != nil {
        return nil, fmt.Errorf("error unmarshaling JSON for SpaceX roadster: %v", err)
    }

	return &roadster, nil
}
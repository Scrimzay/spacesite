package spacexapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpaceXStarlinkAPI struct {
	SpaceTrack SpaceTrack `json:"spaceTrack"`
	Version string `json:"version"`
	Longitude float32 `json:"longitude"`
	Latitude float32 `json:"latitude"`
	HeightKM float32 `json:"height_km"`
	VelocityKM float32 `json:"velocity_km"`
	ID string `json:"id"`
}

type SpaceTrack struct {
	ObjectName string `json:"OBJECT_NAME"`
	Comment string `json:"COMMENT"`
	CreationDate string `json:"CREATION_DATE"`
	ObjectID string `json:"OBJECT_ID"`
	CenterName string `json:"CENTER_NAME"`
	TimeSystem string `json:"TIME_SYSTEM"`
	Epoch string `json:"EPOCH"`
	MeanMotion float32 `json:"MEAN_MOTION"`
	Eccentricity float32 `json:"ECCENTRICITY"`
	Inclination float32 `json:"INCLINATION"`
	ClassificationType string `json:"CLASSIFICATION_TYPE"`
	ObjectType string `json:"OBJECT_TYPE"`
	CountryCode string `json:"COUNTRY_CODE"`
	LaunchDate string `json:"LAUNCH_DATE"`
	Site string `json:"SITE"`
	DecayDate string `json:"DECAY_DATE"`
	Decayed int `json:"DECAYED"`
	TLELine0 string `json:"TLE_LINE0`
	TLELine1 string `json:"TLE_LINE1`
	TLELine2 string `json:"TLE_LINE2`
}

func CallGeneralSpaceXStarlinkInformation() ([]SpaceXStarlinkAPI, error) {
	baseURL := "https://api.spacexdata.com/v4/starlink"

	//send request
	req, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Could not send get to base URL space X starlink: %v", err)
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response for space X starlink: %v", err)
	}

	var satellites []SpaceXStarlinkAPI
	err = json.Unmarshal(body, &satellites)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling json for space X starlink: %v", err)
	}

	// Filter non-decayed satellites and limit to 100 entries
    const maxSatellites = 100
    filteredSatellites := make([]SpaceXStarlinkAPI, 0, maxSatellites) // Pre-allocate with max capacity
    for _, satellite := range satellites {
        if satellite.SpaceTrack.Decayed == 0 {
            filteredSatellites = append(filteredSatellites, satellite)
            if len(filteredSatellites) >= maxSatellites {
                break // Stop once we have 100 non-decayed satellites
            }
        }
    }

    fmt.Printf("Total satellites: %d, Non-decayed satellites: %d", len(satellites), len(filteredSatellites))

    return filteredSatellites, nil
}

func CallSpecificSpaceXStarlinkInformation(id string) (*SpaceXStarlinkAPI, error) {
    baseURL := fmt.Sprintf("https://api.spacexdata.com/v4/starlink/%s", id)
    resp, err := http.Get(baseURL)
    if err != nil {
        return nil, fmt.Errorf("could not send GET to SpaceX starlink API: %v", err)
    }
    defer resp.Body.Close()

    // Read response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response for SpaceX starlink: %v", err)
    }

    var satellite SpaceXStarlinkAPI
    err = json.Unmarshal(body, &satellite)
    if err != nil {
        return nil, fmt.Errorf("error unmarshaling JSON for SpaceX starlink: %v", err)
    }

    return &satellite, nil
}
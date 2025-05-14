package spacexapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpaceXRocketsAPI struct {
	Name          string `json:"name"`
    Type          string `json:"type"`
    Active        bool   `json:"active"`
    Stages        int    `json:"stages"`
    Boosters      int    `json:"boosters"`
    CostPerLaunch int    `json:"cost_per_launch"`
    SuccessRate   int    `json:"success_rate_pct"`
    FirstFlight   string `json:"first_flight"`
    Country       string `json:"country"`
    Company       string `json:"company"`
    Description   string `json:"description"`
	Height        Dimension `json:"height"`
    Diameter      Dimension `json:"diameter"`
    Mass          Weight    `json:"mass"`
    FirstStage    FirstStage `json:"first_stage"`
    SecondStage   SecondStage `json:"second_stage"`
    Engines       Engine     `json:"engines"`
    LandingLegs   LandingLegs `json:"landing_legs"`
    PayloadWeights []PayloadWeight `json:"payload_weights"`
	ID string `json:"id"`
}

// Nested structs for nested JSON objects
type Dimension struct {
    Meters float64 `json:"meters"`
    Feet   float64 `json:"feet"`
}

type Weight struct {
    Kg int `json:"kg"`
    Lb int `json:"lb"`
}

type Thrust struct {
    KN   float64 `json:"kN"`
    Lbf  int     `json:"lbf"`
}

type FirstStage struct {
    ThrustSeaLevel Thrust  `json:"thrust_sea_level"`
    ThrustVacuum   Thrust  `json:"thrust_vacuum"`
    Reusable       bool    `json:"reusable"`
    Engines        int     `json:"engines"`
    FuelAmountTons float64 `json:"fuel_amount_tons"`
    BurnTimeSec    int     `json:"burn_time_sec"`
}

type SecondStage struct {
    Thrust     Thrust      `json:"thrust"`
    Payloads   Payloads    `json:"payloads"`
    Reusable   bool        `json:"reusable"`
    Engines    int         `json:"engines"`
    FuelAmountTons float64 `json:"fuel_amount_tons"`
    BurnTimeSec int         `json:"burn_time_sec"`
}

type Payloads struct {
    CompositeFairing Dimension `json:"composite_fairing"`
    Option1          string    `json:"option_1"`
}

type Engine struct {
    ISP struct {
        SeaLevel int `json:"sea_level"`
        Vacuum   int `json:"vacuum"`
    } `json:"isp"`
    ThrustSeaLevel  Thrust  `json:"thrust_sea_level"`
    ThrustVacuum    Thrust  `json:"thrust_vacuum"`
    Number          int     `json:"number"`
    Type            string  `json:"type"`
    Version         string  `json:"version"`
    Layout          string  `json:"layout"`
    EngineLossMax   int     `json:"engine_loss_max"`
    Propellant1     string  `json:"propellant_1"`
    Propellant2     string  `json:"propellant_2"`
    ThrustToWeight  float64 `json:"thrust_to_weight"`
}

type LandingLegs struct {
    Number   int    `json:"number"`
    Material string `json:"material"`
}

type PayloadWeight struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Kg   int    `json:"kg"`
    Lb   int    `json:"lb"`
}

// RocketListItem is a minimal struct for extracting name and ID from the /rockets endpoint
type RocketListItem struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

func CallGeneralSpaceXRocketsInformation() ([]SpaceXRocketsAPI, error) {
	baseURL := "https://api.spacexdata.com/v4/rockets"

	//send request
	req, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Could not send get to base URL space X rockets: %v", err)
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response for space X rockets: %v", err)
	}

	var rockets []SpaceXRocketsAPI
	err = json.Unmarshal(body, &rockets)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling json for space X rockets: %v", err)
	}

	return rockets, nil
}

func CallSpecificSpaceXRocketsInformation(id string) (*SpaceXRocketsAPI, error) {
    baseURL := fmt.Sprintf("https://api.spacexdata.com/v4/rockets/%s", id)
    resp, err := http.Get(baseURL)
    if err != nil {
        return nil, fmt.Errorf("could not send GET to SpaceX rockets API: %v", err)
    }
    defer resp.Body.Close()

    // Read response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response for SpaceX rockets: %v", err)
    }

    var rocket SpaceXRocketsAPI
    err = json.Unmarshal(body, &rocket)
    if err != nil {
        return nil, fmt.Errorf("error unmarshaling JSON for SpaceX rockets: %v", err)
    }

    return &rocket, nil
}
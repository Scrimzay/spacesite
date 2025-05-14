package spacexapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpaceXDragonsAPI struct {
	Name string `json:"name"`
	Description string `json:"description"`
	FirstFlight string `json:"first_flight"`
	Type string `json:"type"`
	Active bool `json:"active"`
	CrewCapacity int `json:"crew_capacity"`
	SidewallAngleDeg int `json:"sidewall_angle_deg"`
	OrbitDurationYr int `json:"orbit_duration_yr"`
	DryMassKg int `json:"dry_mass_kg"`
	DryMassLb int `json:"dry_mass_lb"`
	Thrusters []Thrusters `json:"thrusters"`
	HeatShield HeatShield `json:"heat_shield"`
	LaunchPayloadMass LaunchPayloadMass `json:"launch_payload_mass"`
	LaunchPayloadVol LaunchPayloadVol `json:"launch_payload_vol"`
	ReturnPayloadMass ReturnPayloadMass `json:"return_payload_mass"`
	ReturnPayloadVol ReturnPayloadVol `json:"return_payload_vol"`
	Diameter Diameter `json:"diameter"`
	PressurizedCapsule PressurizedCapsule `json:"pressurized_capsule"`
	Trunk Trunk `json:"trunk"`
	HeightWithTrunk HeightWithTrunk `json:"height_w_trunk"`
	ID string `json:"id"`
}

type Thrusters struct {
	Type string `json:"type"`
	Amount int `json:"amount"`
	Pods int `json:"pods"`
	Fuel1 string `json:"fuel_1"`
	Fuel2 string `json:"fuel_2"`
	ISP int `json:"isp"`
	ThrustForce ThrustForce `json:"thrust"`
}

type ThrustForce struct {
	KN float32 `json:"kN"`
	LBF int `json:"lbf"`
}

type HeatShield struct {
	Material string `json:"material"`
	SizeMeters float32 `json:"size_meters"`
	TempDegrees int `json:"temp_degrees"`
	DevPartner string `json:"dev_partner"`
}

type LaunchPayloadMass struct {
	KG int `json:"kg"`
	LB int `json:"lb"`
}

type LaunchPayloadVol struct {
	CubicMeters int `json:"cubic_meters"`
	CubicFeet int `json:"cubic_feet"`
}

type ReturnPayloadMass struct {
	KG int `json:"kg"`
	LB int `json:"lb"`
}

type ReturnPayloadVol struct {
	CubicMeters int `json:"cubic_meters"`
	CubicFeet int `json:"cubic_feet"`
}

type PressurizedCapsule struct {
	PayloadVolume PayloadVolume `json:"payload_volume"`
}

type PayloadVolume struct {
	CubicMeters int `json:"cubic_meters"`
	CubicFeet int `json:"cubic_feet"`
}

type Diameter struct {
	Meters float32 `json:"meters"`
	Feet int `json:"feet"`
}

type HeightWithTrunk struct {
	Meters float32 `json:"meters"`
	Feet float32 `json:"feet"`
}

type Trunk struct {
	TrunkVolume TrunkVolume `json:"trunk_volume"`
	Cargo Cargo `json:"cargo"`
}

type TrunkVolume struct {
	CubicMeters int `json:"cubic_meters"`
	CubicFeet int `json:"cubic_feet"`
}

type Cargo struct {
	SolarArray int `json:"solar_array"`
	UnpressurizedCargo bool `json:"unpressurized_cargo"`
}

func CallGeneralSpaceXDragonsInformation() ([]SpaceXDragonsAPI, error) {
	baseURL := "https://api.spacexdata.com/v4/dragons"

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Error sending get request to spacex dragons: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read resp body from spacex dragons: %v", err)
	}

	var dragons []SpaceXDragonsAPI
	err = json.Unmarshal(body, &dragons)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling json resp from spacex dragons: %v", err)
	}

	return dragons, nil
}
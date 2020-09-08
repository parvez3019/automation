package service

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
)

type Appliances struct {
	Location  CorridorLocation
	Appliance ApplianceStateI
}

type AppliancesInfo struct {
	Name             string
	Number           int
	IsSwitchedOn     bool
	PowerConsumption int
	Location         CorridorLocation
}

type CorridorLocation struct {
	FloorNumber    int
	CorridorType   CorridorType
	CorridorNumber int
}

type CreateHotelRequest struct {
	NumberOfFloors       int
	MainCorridorPerFloor int
	SubCorridorPerFloor  int
}

type ToggleApplianceRequest struct {
	AppType  ApplianceType
	TurnOn   bool
	Location CorridorLocation
}

type MovementDetectedEvent struct {
	Movement bool
	Location CorridorLocation
}

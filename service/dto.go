package service

import . "HotelAutomation/model/appliances"

type Appliances struct {
	Location  ApplianceLocation
	Appliance ApplianceStateI
}

type AppliancesInfo struct {
	Name             string
	Number           int
	IsSwitchedOd     bool
	PowerConsumption int
	Location         ApplianceLocation
}

type ApplianceLocation struct {
	FloorNumber    int
	CorridorType   string
	CorridorNumber int
}

type CreateHotelRequest struct {
	NumberOfFloors       int
	MainCorridorPerFloor int
	SubCorridorPerFloor  int
}

type ToggleApplianceRequest struct {
	ApplianceType   ApplianceType
	ApplianceNumber int
	SwitchOn        bool
	Location        ApplianceLocation
}

package service

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
)

func mapApplianceToApplianceInfo(appliances []Appliances) []AppliancesInfo {
	appliancesInfo := make([]AppliancesInfo, 0)
	for _, a := range appliances {
		appliancesInfo = append(appliancesInfo, AppliancesInfo{
			Name:             a.Appliance.GetType(),
			Number:           a.Appliance.GetId(),
			IsSwitchedOd:     a.Appliance.IsOn(),
			PowerConsumption: a.Appliance.GetPowerConsumption(),
			Location:         a.Location,
		})
	}
	return appliancesInfo
}

func mapToAppliances(floors []*Floor) []Appliances {
	applianceStates := make([]Appliances, 0)
	for _, floor := range floors {
		applianceStates = append(applianceStates, mapCorridorLocation(floor, MAIN)...)
		applianceStates = append(applianceStates, mapCorridorLocation(floor, SUB)...)
	}
	return applianceStates
}

func mapCorridorLocation(floor *Floor, corridorType CorridorType) []Appliances {
	applianceStates := make([]Appliances, 0)
	for _, corridor := range floor.GetCorridors(corridorType) {
		applianceLocation := getApplianceLocation(floor, corridor)
		applianceStates = mapLightBulb(corridor, applianceStates, applianceLocation)
		applianceStates = mapAC(corridor, applianceStates, applianceLocation)
	}
	return applianceStates
}

func mapAC(corridor *Corridor, applianceStates []Appliances, applianceLocation CorridorLocation) []Appliances {
	for _, ac := range corridor.GetAirConditioners() {
		applianceStates = append(applianceStates, mapToApplianceState(ac, applianceLocation))
	}
	return applianceStates
}

func mapLightBulb(corridor *Corridor, applianceStates []Appliances, applianceLocation CorridorLocation) []Appliances {
	for _, bulb := range corridor.GetLightBulbs() {
		applianceStates = append(applianceStates, mapToApplianceState(bulb, applianceLocation))
	}
	return applianceStates
}

func getApplianceLocation(floor *Floor, corridor *Corridor) CorridorLocation {
	return CorridorLocation{
		FloorNumber:    floor.GetLevel(),
		CorridorType:   corridor.GetType(),
		CorridorNumber: corridor.GetId(),
	}
}

func mapToApplianceState(appliance ApplianceStateI, location CorridorLocation) Appliances {
	return Appliances{
		Location:  location,
		Appliance: appliance,
	}
}

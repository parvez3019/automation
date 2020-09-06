package service

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
)

func mapToApplianceStateDto(floors []*Floor) []ApplianceInfo {
	applianceStates := make([]ApplianceInfo, 0)
	for _, floor := range floors {
		applianceStates = append(applianceStates, mapCorridorLocation(floor, MAIN)...)
		applianceStates = append(applianceStates, mapCorridorLocation(floor, SUB)...)
	}
	return applianceStates
}

func mapCorridorLocation(floor *Floor, corridorType CorridorType) []ApplianceInfo {
	applianceStates := make([]ApplianceInfo, 0)
	for _, corridor := range floor.GetCorridors(corridorType) {
		applianceLocation := getApplianceLocation(floor, corridor)
		applianceStates = mapLightBulb(corridor, applianceStates, applianceLocation)
		applianceStates = mapAC(corridor, applianceStates, applianceLocation)
	}
	return applianceStates
}

func mapAC(corridor *Corridor, applianceStates []ApplianceInfo, applianceLocation ApplianceLocation) []ApplianceInfo {
	for _, ac := range corridor.GetAirConditioners() {
		applianceStates = append(applianceStates, mapToApplianceState(AC, ac, applianceLocation))
	}
	return applianceStates
}

func mapLightBulb(corridor *Corridor, applianceStates []ApplianceInfo, applianceLocation ApplianceLocation) []ApplianceInfo {
	for _, bulb := range corridor.GetLightBulbs() {
		applianceStates = append(applianceStates, mapToApplianceState(LIGHT, bulb, applianceLocation))
	}
	return applianceStates
}

func getApplianceLocation(floor *Floor, corridor *Corridor) ApplianceLocation {
	return ApplianceLocation{
		floorNumber:    floor.GetLevel(),
		corridorType:   corridor.GetTypeAsString(),
		corridorNumber: corridor.GetId(),
	}
}

func mapToApplianceState(applianceType ApplianceType, appliance ApplianceI, location ApplianceLocation) ApplianceInfo {
	return ApplianceInfo{
		Name:             string(applianceType),
		Number:           appliance.GetId(),
		IsSwitchedOd:     appliance.IsOn(),
		PowerConsumption: appliance.GetPowerConsumption(),
		Location:         location,
	}
}
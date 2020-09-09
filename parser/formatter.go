package parser

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	. "HotelAutomation/service"
	"fmt"
	"sort"
	"strings"
)

type Formatter struct {
}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (o *Formatter) ApplianceInfoToString(appliances []AppliancesInfo) string {
	var str strings.Builder
	floors := getSortedFloors(appliances)
	for _, floor := range floors {
		str.WriteString(fmt.Sprintf("Floor %d\n", floor))
		str.WriteString(printCorridor(floor, MAIN, appliances))
		str.WriteString(printCorridor(floor, SUB, appliances))
	}
	return str.String()
}

func printCorridor(floor int, cType CorridorType, appliances []AppliancesInfo) string {
	var str strings.Builder
	corridors := getSortedCorridors(floor, cType, appliances)
	for _, corridor := range corridors {
		str.WriteString(fmt.Sprintf("%s Corridor %d ", cType, corridor))
		str.WriteString(printAppliance(appliances, floor, cType, corridor, LIGHT))
		str.WriteString(" : ")
		str.WriteString(printAppliance(appliances, floor, cType, corridor, AC))
		str.WriteString("\n")
	}
	return str.String()
}

func printAppliance(info []AppliancesInfo, floor int, cType CorridorType, corridor int, aType ApplianceType) string {
	var str strings.Builder
	appliances := getSortedAppliance(info, floor, cType, corridor, aType)
	for number, state := range appliances {
		str.WriteString(fmt.Sprintf("%s %d : %s ", string(aType), number, appToggleString(state)))
	}
	return str.String()
}

func getSortedAppliance(info []AppliancesInfo, floor int, cType CorridorType, corridor int, aType ApplianceType) map[int]bool {
	uniqueAppliances := make(map[int]bool, 0)
	for _, a := range info {
		l := a.Location
		if l.FloorNumber == floor && l.CorridorType == cType && l.CorridorNumber == corridor && a.Name == string(aType) {
			uniqueAppliances[a.Number] = a.IsSwitchedOn
		}
	}
	return uniqueAppliances
}

func getSortedCorridors(floor int, cType CorridorType, appliances []AppliancesInfo) []int {
	uniqueCorridors := make(map[int]int, 0)
	for _, appliance := range appliances {
		if appliance.Location.FloorNumber == floor && appliance.Location.CorridorType == cType {
			uniqueCorridors[appliance.Location.CorridorNumber] = 1
		}
	}
	return mapOfKeysToSlice(uniqueCorridors)
}

func getSortedFloors(appliances []AppliancesInfo) []int {
	uniqueFloors := make(map[int]int, 0)
	for _, appliance := range appliances {
		uniqueFloors[appliance.Location.FloorNumber] = 1
	}
	return mapOfKeysToSlice(uniqueFloors)
}

func mapOfKeysToSlice(uniqueFloors map[int]int) []int {
	floorsAsList := make([]int, 0)
	for k, _ := range uniqueFloors {
		floorsAsList = append(floorsAsList, k)
	}
	sort.Ints(floorsAsList)
	return floorsAsList
}

func appToggleString(on bool) string {
	if on {
		return "ON"
	}
	return "OFF"
}

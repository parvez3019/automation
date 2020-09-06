package service

type ApplianceInfo struct {
	Name             string
	Number           int
	IsSwitchedOd     bool
	PowerConsumption int
	Location ApplianceLocation
}

type ApplianceLocation struct {
	floorNumber    int
	corridorType   string
	corridorNumber int
}

type CreateHotelRequest struct {
	NumberOfFloors       int
	MainCorridorPerFloor int
	SubCorridorPerFloor  int
}

type ApplianceType string

const (
	LIGHT ApplianceType = "Light"
	AC    ApplianceType = "AC"
)

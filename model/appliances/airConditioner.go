package appliances

type AirConditioner struct {
	*Appliance
}

func NewAirConditioner(id int, powerConsumption int) *AirConditioner {
	return &AirConditioner{NewAppliance(id, powerConsumption, AC)}
}

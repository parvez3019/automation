package appliances


type LightBulb struct {
	*Appliance
}

func NewLightBulb(id, powerConsumption int) *LightBulb {
	return &LightBulb{NewAppliance(id, powerConsumption, LIGHT)}
}

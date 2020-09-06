package appliances

type ApplianceI interface {
	TurnOn()
	TurnOff()
	IsOn() bool
	GetId() int
	GetPowerConsumption() int
}

type Appliance struct {
	id               int
	on               bool
	powerConsumption int
}

func NewAppliance(id int, powerConsumption int) *Appliance {
	return &Appliance{id: id, powerConsumption: powerConsumption,}
}

func (appliance *Appliance) TurnOn() {
	appliance.on = true
}

func (appliance *Appliance) TurnOff() {
	appliance.on = false
}

func (appliance *Appliance) IsOn() bool {
	return appliance.on
}

func (appliance *Appliance) GetId() int {
	return appliance.id
}

func (appliance *Appliance) GetPowerConsumption() int {
	return appliance.powerConsumption
}



























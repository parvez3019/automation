package appliances

type ApplianceI interface {
	turnOn()
	turnOff()
	isOn() bool
	getId() int
	getPowerConsumption() int
}

type Appliance struct {
	id               int
	on               bool
	powerConsumption int
}

func NewAppliance(id int, powerConsumption int) *Appliance {
	return &Appliance{id: id, powerConsumption: powerConsumption,}
}

func (appliance *Appliance) turnOn() {
	appliance.on = true
}

func (appliance *Appliance) turnOff() {
	appliance.on = false
}

func (appliance *Appliance) isOn() bool {
	return appliance.on
}

func (appliance *Appliance) getId() int {
	return appliance.id
}

func (appliance *Appliance) getPowerConsumption() int {
	return appliance.powerConsumption
}



























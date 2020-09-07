package appliances

type ApplianceInfoI interface {
	IsOn() bool
	GetId() int
	GetType() string
	GetPowerConsumption() int
}

type ApplianceStateI interface {
	IsOn() bool
	GetId() int
	GetType() string
	GetPowerConsumption() int
	SetSwitchedOn(bool)
}

type Appliance struct {
	id               int
	on               bool
	powerConsumption int
	applianceType    ApplianceType
}

func NewAppliance(id int, powerConsumption int, applianceType ApplianceType) *Appliance {
	return &Appliance{id: id, powerConsumption: powerConsumption, applianceType: applianceType}
}

func (appliance *Appliance) SetSwitchedOn(value bool) {
	appliance.on = value
}

func (appliance *Appliance) IsOn() bool {
	return appliance.on
}

func (appliance *Appliance) GetId() int {
	return appliance.id
}

func (appliance *Appliance) GetType() string {
	return string(appliance.applianceType)
}

func (appliance *Appliance) GetPowerConsumption() int {
	return appliance.powerConsumption
}

type ApplianceType string

const (
	LIGHT ApplianceType = "Light"
	AC    ApplianceType = "AC"
)

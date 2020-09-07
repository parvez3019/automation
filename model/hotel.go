package model

type Hotel struct {
	floors []*Floor
}

func NewHotel() *Hotel {
	return &Hotel{}
}

func (h *Hotel) AddFloors(floor []*Floor) *Hotel {
	h.floors = append(h.floors, floor...)
	return h
}

func (h *Hotel) GetFloors() []*Floor {
	return h.floors
}

func (h *Hotel) GetCorridors(corridorType CorridorType) []*Corridor {
	corridors := make([]*Corridor, 0)
	for _, floor := range h.floors {
		corridors = append(corridors, floor.GetCorridors(corridorType)...)
	}
	return corridors
}

func (h *Hotel) GetAllCorridors() []*Corridor {
	return append(h.GetCorridors(MAIN), h.GetCorridors(SUB)...)
}

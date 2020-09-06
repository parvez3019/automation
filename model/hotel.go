package model

type Hotel struct {
	floors []*Floor
}

func NewHotel() *Hotel {
	return &Hotel{}
}

func (h *Hotel) addFloors(floor []*Floor) *Hotel {
	h.floors = append(h.floors, floor...)
	return h
}

func (h *Hotel) getFloors() []*Floor {
	return h.floors
}

func (h *Hotel) getCorridors() []*Corridor {
	corridors := make([]*Corridor, 0)
	for _, floor := range h.floors {
		corridors = append(corridors, floor.getCorridors(MAIN)...)
		corridors = append(corridors, floor.getCorridors(SUB)...)
	}
	return corridors
}

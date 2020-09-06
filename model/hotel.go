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

func(h *Hotel) getFloors() []*Floor {
	return h.floors
}



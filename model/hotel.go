package model

type hotel struct {
	floors []*floor
}

func NewHotel() *hotel {
	return &hotel{}
}

func (h *hotel) addFloor(floor *floor) *hotel{
	h.floors = append(h.floors, floor)
	return h
}

func(h *hotel) getFloors() []*floor {
	return h.floors
}



package model

type floor struct {
	level         int
	mainCorridors []*Corridor
	subCorridors  []*Corridor
}

func NewFloor(level int) *floor {
	return &floor{level: level}
}

func (f *floor) getLevel() int {
	return f.level
}

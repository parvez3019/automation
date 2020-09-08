package model

type Floor struct {
	level         int
	mainCorridors []*Corridor
	subCorridors  []*Corridor
}

func NewFloor(level int) *Floor {
	return &Floor{level: level}
}

func (f *Floor) GetLevel() int {
	return f.level
}

func (f *Floor) GetCorridors(cType CorridorType) []*Corridor {
	if cType == MAIN {
		return f.mainCorridors
	}
	if cType == SUB {
		return f.subCorridors
	}
	return []*Corridor{}
}

func (f *Floor) AddCorridors(corridors []*Corridor) *Floor {
	for _, corridor := range corridors {
		if corridor.GetType() == MAIN {
			f.mainCorridors = append(f.mainCorridors, corridor)
		}
		if corridor.GetType() == SUB {
			f.subCorridors = append(f.subCorridors, corridor)
		}
	}
	return f
}

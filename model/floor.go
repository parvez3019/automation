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

func (f *Floor) getCorridors(cType Type) []*Corridor {
	if cType == MAIN {
		return f.mainCorridors
	}
	return f.subCorridors
}

func (f *Floor) addCorridors(c []*Corridor, cType Type) *Floor {
	if cType == MAIN {
		f.mainCorridors = append(f.mainCorridors, c...)
	}
	if cType == SUB {
		f.subCorridors = append(f.subCorridors, c...)
	}
	return f
}
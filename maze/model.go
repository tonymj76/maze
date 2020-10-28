package maze

type Coordinate struct {
	X, Y string
}

type Path struct {
	Distance float32
}

type Spot struct {
	Name         string
	Coordinates  Coordinate
	AmountOfGold int32 `json:"amount_of_gold"`
}

type Quadrant struct {
	TopRight    Coordinate `json:"top_right"`
	TopLeft     Coordinate `json:"top_left"`
	BottomRight Coordinate `json:"bottom_right"`
	BottomLeft  Coordinate `json:"bottom_left"`
}

type Maze struct {
	Path     Path
	Spots    []Spot
	Quadrant Quadrant
}

type Mazes []*Maze

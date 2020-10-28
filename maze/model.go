package maze

import "github.com/google/uuid"

type Coordinate struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type Spot struct {
	Name         string     `json:"name"`
	Coordinates  Coordinate `json:"coordinates"`
	AmountOfGold int32      `json:"amount_of_gold"`
}

type Quadrant struct {
	TopRight    Coordinate `json:"top_right,omitempty"`
	TopLeft     Coordinate `json:"top_left,omitempty"`
	BottomRight Coordinate `json:"bottom_right,omitempty"`
	BottomLeft  Coordinate `json:"bottom_left,omitempty"`
}

type Maze struct {
	ID           uuid.UUID `json:"id"`
	PathDistance float32   `json:"path_distance"`
	Spots        []Spot    `json:"spots"`
	Quadrant     Quadrant  `json:"quandrant"`
}

type Mazes []*Maze

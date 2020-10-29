package maze

import "github.com/globalsign/mgo/bson"

// Coordinate points
type Coordinate struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

// Spot position
type Spot struct {
	Name         string     `json:"name"`
	Coordinates  Coordinate `json:"coordinates"`
	AmountOfGold int32      `json:"amount_of_gold"`
}

// Quadrant _
type Quadrant struct {
	TopRight    Coordinate `json:"top_right,omitempty"`
	TopLeft     Coordinate `json:"top_left,omitempty"`
	BottomRight Coordinate `json:"bottom_right,omitempty"`
	BottomLeft  Coordinate `json:"bottom_left,omitempty"`
}

// Maze _
type Maze struct {
	ID           bson.ObjectId `bson:"_id" json:"id,omitempty"`
	PathDistance float32       `json:"path_distance"`
	Spots        []Spot        `json:"spots"`
	Quadrant     Quadrant      `json:"quandrant"`
}

//Mazes list of maze
type Mazes []*Maze

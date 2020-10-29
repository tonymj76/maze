package maze

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Repository _
type Repository interface {
	GetByMazeID(string) (*Maze, error)
	Create(*Maze) error
	Update(string, *Maze) error
	Get() (Mazes, error)
	Close()
	Delete(string) error
}

//MazeRepository struct
type MazeRepository struct {
	Session *mgo.Session
}

const (
	databaseName   = "mazes"
	collectionName = "mazes"
)

// GetByMazeID get a maze id
func (r *MazeRepository) GetByMazeID(mazeID string) (*Maze, error) {
	var maze *Maze
	if ok := bson.IsObjectIdHex(mazeID); !ok {
		return nil, fmt.Errorf("%s is not a vaild _id", mazeID)
	}
	objID := bson.ObjectIdHex(mazeID)
	fmt.Println(objID)
	err := r.collection().Find(bson.M{
		"_id": objID,
	}).One(&maze)
	if err != nil {
		return nil, err
	}
	return maze, nil
}

// Get get all maze entries in he db
func (r *MazeRepository) Get() (Mazes, error) {
	var mazes Mazes
	err := r.collection().Find(nil).All(&mazes)
	if err != nil {
		return nil, err
	}
	return mazes, nil
}

// Update a particular maze object
func (r *MazeRepository) Update(id string, updateMaze *Maze) error {
	if ok := bson.IsObjectIdHex(id); !ok {
		return fmt.Errorf("%s is not a vaild _id", id)
	}
	objID := bson.ObjectIdHex(id)
	return r.collection().UpdateId(objID, updateMaze)
}

// Create maze
func (r *MazeRepository) Create(maze *Maze) error {
	return r.collection().Insert(maze)
}

//Delete maze
func (r *MazeRepository) Delete(mazeID string) error {
	if ok := bson.IsObjectIdHex(mazeID); !ok {
		return fmt.Errorf("%s is not a vaild _id", mazeID)
	}
	objID := bson.ObjectIdHex(mazeID)
	return r.collection().Remove(bson.M{"_id": objID})
}

// Close the session created
func (r *MazeRepository) Close() {
	r.Session.Close()
}

func (r *MazeRepository) collection() *mgo.Collection {
	return r.Session.DB(databaseName).C(collectionName)
}

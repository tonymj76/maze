package maze

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Repository interface {
	GetByMazeID(string) (*Maze, error)
	Create(*Maze) error
	Update(string, *Maze) error
	Get() (Mazes, error)
	Close()
	Delete(string) error
}
type MazeRepository struct {
	Session *mgo.Session
}

const (
	databaseName   = "mazes"
	collectionName = "mazes"
)

func (r *MazeRepository) GetByMazeID(mazeID string) (*Maze, error) {
	var maze *Maze
	err := r.collection().Find(bson.M{
		"_id": mazeID,
	}).One(&maze)
	if err != nil {
		return nil, err
	}
	return maze, nil
}
func (r *MazeRepository) Get() (Mazes, error) {
	var mazes Mazes
	err := r.collection().Find(nil).All(&mazes)
	if err != nil {
		return nil, err
	}
	return mazes, nil
}
func (r *MazeRepository) Update(id string, updateMaze *Maze) error {
	return r.collection().UpdateId(id, updateMaze)
}
func (r *MazeRepository) Create(maze *Maze) error {
	return r.collection().Insert(maze)
}

func (r *MazeRepository) Delete(mazeID string) error {
	return r.collection().Remove(bson.M{"_id": mazeID})
}

func (repo *MazeRepository) Close() {
	repo.Session.Close()
}

func (repo *MazeRepository) collection() *mgo.Collection {
	return repo.Session.DB(databaseName).C(collectionName)
}

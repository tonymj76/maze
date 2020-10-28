package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/tonymj76/maze/maze"
)

type Service struct {
	Session *mgo.Session
}

func (s *Service) GetRepo() maze.Repository {
	return &maze.MazeRepository{Session: s.Session.Clone()}
}

func (s *Service) CreateHandler(c *gin.Context) {
	var maze *maze.Maze
	if err := c.ShouldBindJSON(&maze); err != nil {
		JSON(c, "failed to bind request", http.StatusBadRequest, nil, err)
	}
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Create(maze); err != nil {
		JSON(c, "failed to bind request", http.StatusBadRequest, nil, err)
	}
	JSON(c, "successful", http.StatusCreated, maze, nil)
}

func (s *Service) UpdateHandler(c *gin.Context) {
	var maze *maze.Maze
	mazeID := c.Query("maze_id")
	if err := c.ShouldBindJSON(&maze); err != nil {
		JSON(c, "failed to bind request", http.StatusBadRequest, nil, err)
	}
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Update(mazeID, maze); err != nil {
		JSON(c, "failed to bind request", http.StatusBadRequest, nil, err)
	}
	JSON(c, "successful", http.StatusCreated, maze, nil)
}

func (s *Service) GetIDHandler(c *gin.Context) {
	mazeID := c.Query("maze_id")
	repo := s.GetRepo()
	defer repo.Close()
	maze, err := repo.GetByMazeID(mazeID)
	if err != nil {
		JSON(c, "failed to bind request", http.StatusBadRequest, nil, err)
	}
	JSON(c, "successful", http.StatusOK, maze, nil)
}

func (s *Service) DeleteIDHandler(c *gin.Context) {
	mazeID := c.Query("maze_id")
	repo := s.GetRepo()
	defer repo.Close()
	err := repo.Delete(mazeID)
	if err != nil {
		JSON(c, "failed to bind request", http.StatusBadRequest, nil, err)
	}
	JSON(c, "successful", http.StatusOK, nil, nil)
}

func (s *Service) GetHandler(c *gin.Context) {
	repo := s.GetRepo()
	defer repo.Close()
	mazes, err := repo.Get()
	if err != nil {
		JSON(c, "failed to bind request", http.StatusBadRequest, nil, err)
	}
	JSON(c, "successful", http.StatusOK, mazes, nil)
}

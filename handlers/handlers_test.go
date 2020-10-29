package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tonymj76/maze/mongodb"
)

type router struct {
	R *gin.Engine
}

func (r *router) performRequest(method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.R.ServeHTTP(w, req)
	return w
}

func TestMain(m *testing.M) {
	code := 0
	defer func() {
		os.Exit(code)
	}()
	url := os.Getenv("URL")
	if url == "" {
		os.Exit(code)
	}
	session, err := mongodb.NewSeasion(url)
	if err != nil {
		log.Fatalf("connection error %v", err)
	}
	s := Service{Session: session}
	r := SetupRouter(s)
	_ = router{R: r}
	code = m.Run()
}
func (r *router) TestGetHandler(t *testing.T) {
	w := r.performRequest("Get", "/api/v1/maze")
	assert.Equal(t, http.StatusOK, w.Code)
}

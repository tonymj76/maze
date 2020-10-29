package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonymj76/maze/mongodb"
)

// type router struct {
// 	R *gin.Engine
// }
var url = "mongodb://127.0.0.1:27017"

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// func TestMain(m *testing.M) {
// 	code := 0
// 	defer func() {
// 		os.Exit(code)
// 	}()
// 	url := os.Getenv("URL")
// 	if url == "" {
// 		os.Exit(code)
// 	}
// 	session, err := mongodb.NewSeasion(url)
// 	if err != nil {
// 		log.Fatalf("connection error %v", err)
// 	}
// 	s := Service{Session: session}
// 	r := SetupRouter(s)
// 	_ = router{R: r}
// 	code = m.Run()
// }
func TestGetHandler(t *testing.T) {
	session, _ := mongodb.NewSeasion(url)
	r := SetupRouter(Service{Session: session})
	w := performRequest(r, "GET", "/api/v1/maze/", nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostMaze(t *testing.T) {
	data := `{
    "path_distance":3454.5,
    "spots":[
        {
            "name":"bottom right quandrate",
            "coordinates": {"x": -3, "y":-8},
            "amount_of_gold": 366
        }
    ],
    "quandrant": {
        "top_left":{
            "x":-6, "y":6
        }
    }
}`
	session, _ := mongodb.NewSeasion(url)
	r := SetupRouter(Service{Session: session})
	w := performRequest(r, "POST", "/api/v1/maze/", bytes.NewBuffer([]byte(data)))
	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

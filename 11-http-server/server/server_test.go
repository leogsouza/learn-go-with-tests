package server

import (
	"net/http"
	"net/http/httptest"
	"fmt"
	"testing"
)

func TestGetPlayers(t *testing.T) {
    
    store := StubPlayerStore{
        map[string]int {
            "Pepper": 20,
            "Floyd": 10,
        },
    }
    
    server := &PlayerServer{&store}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})
	
	t.Run("returns Floyd's score", func(t *testing.T) {
	    request := newGetScoreRequest("Floyd")
	    response := httptest.NewRecorder()
	    
	    server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	    
	})
	
	t.Run("returning 404 on missing players", func(t *testing.T) {
	    request := newGetScoreRequest("Apollo")
	    response := httptest.NewRecorder()
	    
	    server.ServeHTTP(response, request)
	    
	    assertStatus(t, response.Code, http.StatusNotFound)
	    
	})
}

func TestStoreWins(t *testing.T) {
    store := StubPlayerStore{
        map[string]int{},
    }
    
    server := &PlayerServer{&store}
    
    t.Run("it returns accepted on POST", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
        response := httptest.NewRecorder()
        
        server.ServeHTTP(response, request)
        
        assertStatus(t, response.Code, http.StatusAccepted)
    })
}

func newGetScoreRequest(name string) *http.Request {
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
    return req
}

func assertResponseBody(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("response body is wrong, got '%s', want '%s'", got, want)
    }
}

func assertStatus(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("did not get correct status, got %d, want %d", got, want)
    }
}
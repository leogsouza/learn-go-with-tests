package server

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
    GetPlayerScore(name string) int
}

type StubPlayerStore struct {
    scores map[string]int
}

type PlayerServer struct {
    Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    
    switch r.Method {
        case http.MethodPost:
            p.processWin(w)
        case http.MethodGet:
            p.showScore(w, r)
    }
	
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Path[len("/players/"):]
    
    score := p.Store.GetPlayerScore(player)
    
    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }
    
    fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
    w.WriteHeader(http.StatusAccepted)
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    
    return score
}

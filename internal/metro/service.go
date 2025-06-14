package metro

import (
	"metro/internal/users"
	"net/http"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s *Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /", s.homePage)
}

func (s *Service) homePage(w http.ResponseWriter, r *http.Request) {
	user := users.GetUser(r)
	stations := getStations()
	home(user, stations).Render(r.Context(), w)
}

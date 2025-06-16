package metro

import (
	"metro/internal/server"
	"metro/internal/users"
	"net/http"
	"strconv"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s *Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /", s.homePage)
	mux.HandleFunc("GET /stations/{id}", s.stationPage)
}

func (s *Service) homePage(w http.ResponseWriter, r *http.Request) {
	user := users.GetUser(r)
	stations := getStations()
	home(user, stations).Render(r.Context(), w)
}

func (s *Service) stationPage(w http.ResponseWriter, r *http.Request) {
	user := users.GetUser(r)
    idString := r.PathValue("id")
    id, err := strconv.Atoi(idString)
    if err != nil {
        http.Error(w, "Incorrect id", http.StatusBadRequest)
        return
    }
    trains, err := getTrains(id)
    if err != nil {
        server.ServerError(w)
        panic(err)
        return
    }
    station(user, id, trains).Render(r.Context(), w)
}

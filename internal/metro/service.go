package metro

import (
	"context"
	"metro/internal/server"
	"metro/internal/users"
	"net/http"
	"strconv"
	"time"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s *Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /", s.homePage)
	mux.HandleFunc("GET /stations/{id}", s.stationPage)
	mux.HandleFunc("GET /stations/{id}/sse", s.stationPageSSE)
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

func (s *Service) stationPageSSE(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
    timerDuration := 3 * time.Second
    timer := time.NewTimer(timerDuration)
	for {
		select {
		case <-timer.C:
			trains, _ := getTrains(id)
			w.Write([]byte("event:trains\ndata: "))
			trainList(trains).Render(context.Background(), w)
			w.Write([]byte("\n\n"))
            timer = time.NewTimer(timerDuration)
		case <-r.Context().Done():
			break
		}
	}
}

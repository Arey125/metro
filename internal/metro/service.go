package metro

import (
	"context"
	"metro/internal/server"
	"metro/internal/users"
	"net/http"
	"strconv"
	"time"
)

type Service struct{
    schema Schema
}

func NewService() Service {
    schema := NewSchema()
	return Service{
        schema: schema,
    }
}

func (s *Service) Register(mux *http.ServeMux) {
    middleware := func (handler http.HandlerFunc) http.Handler {
		return users.OnlyWithPermission(
			http.HandlerFunc(handler),
			users.PermissonCanUseApplication,
		)
    }
	mux.HandleFunc("GET /", s.homePage)
	mux.Handle("GET /stations/{id}", middleware(s.stationPage))
	mux.Handle("GET /stations/{id}/sse", middleware(s.stationPageSSE))
}

func (s *Service) homePage(w http.ResponseWriter, r *http.Request) {
	user := users.GetUser(r)
	stations := s.schema.stations
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

    station := s.schema.getStation(id)
    if station == nil {
        http.Error(w,"Station not found", http.StatusNotFound)
    }

	trains, err := getTrains(id)
	if err != nil {
		server.ServerError(w, err)
		return
	}

	stationPageTemplate(user, *station, trains).Render(r.Context(), w)
}

func (s *Service) stationPageSSE(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.WriteHeader(http.StatusOK)
	timerDuration := 3 * time.Second
	timer := time.NewTimer(timerDuration)
	for {
		select {
		case <-r.Context().Done():
			return
		case <-timer.C:
			trains, _ := getTrains(id)
			w.Write([]byte("event:trains\ndata: "))
			trainList(trains).Render(context.Background(), w)
			w.Write([]byte("\n\n"))
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			timer = time.NewTimer(timerDuration)
		}
	}
}

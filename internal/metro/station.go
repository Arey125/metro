package metro

import (
	"metro/internal/server"
	"metro/internal/users"
	"net/http"
	"strconv"
	"time"
)

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
		http.Error(w, "Station not found", http.StatusNotFound)
	}

	trains, err := getTrains(id)
	if err != nil {
		server.ServerError(w, err)
		return
	}

	s.stationPageTemplate(user, *station, trains).Render(r.Context(), w)
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
			s.trainList(trains).Render(r.Context(), w)
			w.Write([]byte("\n\n"))
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			timer = time.NewTimer(timerDuration)
		}
	}
}

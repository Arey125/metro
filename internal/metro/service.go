package metro

import (
	"metro/internal/users"
	"net/http"
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

func (s *Service) Start() {
	go s.DataCollectionWorker()
}

func (s *Service) homePage(w http.ResponseWriter, r *http.Request) {
	user := users.GetUser(r)
	stations := s.schema.stations
	home(user, stations).Render(r.Context(), w)
}

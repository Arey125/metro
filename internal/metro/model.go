package metro

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"time"
)

type StationSnapshot struct {
	Id        int
	StationId int
	CreatedAt time.Time
	Response  []byte
}

type Model struct {
	db *sql.DB
}

func NewModel(db *sql.DB) Model {
	return Model{db}
}

func (m *Model) AddStationSnapshot(snapshot StationSnapshot) error {
	_, err := sq.Insert("station_snapshots").
		Columns("station_id", "created_at", "response").
		Values(snapshot.StationId, snapshot.CreatedAt, snapshot.Response).
		RunWith(m.db).
		Exec()

	return err
}

CREATE TABLE station_snapshots (
    id INTEGER PRIMARY KEY,
    station_id INTEGER,
    created_at TEXT,
    response BLOB
);

CREATE INDEX station_snapshots_station_id_idx ON station_snapshots(station_id);

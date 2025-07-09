select s.id, key, station_id, created_at, value->>'$[0].id' as train_id
from station_snapshots as s, json_each((decode(s.response)::json).data)
where train_id = '765_65359-765_65360'

package metro

import (
	"encoding/json"
	"os"
)

type Name string

func (name *Name) UnmarshalJSON(data []byte) error {
	temp := struct {
		Ru string `json:"ru"`
	}{}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	*name = Name(temp.Ru)
	return nil
}

type Station struct {
	Id   int  `json:"id"`
	Name Name `json:"name"`
}

type SchemaData struct {
	Stations []Station `json:"stations"`
}

type Schema struct {
	stationMap map[int]*Station
	stations   []Station
}

func NewSchema() Schema {
	fileBytes, err := os.ReadFile("./assets/schema.json")
	if err != nil {
		panic(err)
	}
	data := struct {
		Data SchemaData `json:"data"`
	}{}
	json.Unmarshal(fileBytes, &data)

	stations := data.Data.Stations
    stationMap := make(map[int]*Station)
    for i, station := range stations {
        stationMap[station.Id] = &stations[i]
    }

    return Schema{
        stations: stations,
        stationMap: stationMap,
    }
}

func (s *Schema) getStation(id int) *Station {
    return s.stationMap[id]
}

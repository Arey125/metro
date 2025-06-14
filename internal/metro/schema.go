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

type Schema struct {
	Stations []Station `json:"stations"`
}

func getStations() []Station {
	fileBytes, err := os.ReadFile("./assets/schema.json")
	if err != nil {
		panic(err)
	}
	data := struct {
		Data Schema `json:"data"`
	}{}
	json.Unmarshal(fileBytes, &data)

	return data.Data.Stations
}

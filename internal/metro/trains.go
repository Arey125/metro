package metro

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Train struct {
	Id string `json:"id"`
    ArrivalTime int `json:"arrivalTime"`
}

type Trains map[int][]Train

func getTrains(stationId int) (Trains, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://prodapp.mosmetro.ru/api/stations/v2/%d/wagons", stationId),
	)
	if err != nil {
		return nil, err
	}
    respObj := struct {
        Data Trains `json:"data"`
    }{}
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Data, nil
}

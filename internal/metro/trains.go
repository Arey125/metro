package metro

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Train struct {
	Id string `json:"id"`
    ArrivalTime int `json:"arrivalTime"`
}

type Trains map[int][]Train

func getTrainsJson(stationId int) ([]byte, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://prodapp.mosmetro.ru/api/stations/v2/%d/wagons", stationId),
	)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}

func getTrains(stationId int) (Trains, error) {
	trainsJson, err := getTrainsJson(stationId)
	if err != nil {
		return nil, err
	}
	respObj := struct {
        Data Trains `json:"data"`
    }{}
	err = json.Unmarshal(trainsJson, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Data, nil
}

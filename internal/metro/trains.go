package metro

import (
	"bytes"
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
    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("%s\n", string(body))
    decoder := json.NewDecoder(bytes.NewReader(body))
    err = decoder.Decode(&respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Data, nil
}

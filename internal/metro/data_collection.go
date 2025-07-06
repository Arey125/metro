package metro

import (
	"fmt"
	"time"
)

var stationIds []int = []int{
	1,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	26,
	33,
	34,
	35,
	36,
	37,
	38,
	39,
	638,
}

func (s *Service) DataCollectionWorker() {
	tickerDuration := time.Duration(s.config.DataCollectionIntervalMs) * time.Millisecond
	ticker := time.NewTicker(tickerDuration)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("At %s", time.Now().Format("Jan 2 2006 15:04:05\n"))
			for _, id := range stationIds {
				station := s.schema.getStation(id)
				trains, _ := getTrainsJson(id)
				fmt.Printf("%s: %s\n", station.Name, trains)
			}
		}
	}
}

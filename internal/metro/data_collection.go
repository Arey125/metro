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

func (s *Service) saveStationSnapshots() {
    errorCount := 0
    now := time.Now()
	for _, id := range stationIds {
		trainsJson, err := getTrainsJson(id)
		if err != nil {
            fmt.Printf("Error occurred while getting station snapshot: %s\n", err)
            errorCount++
			continue
		}
        err = s.model.AddStationSnapshot(StationSnapshot{
			StationId: id,
			CreatedAt: time.Now(),
			Response: trainsJson,
		})
        if err != nil {
            fmt.Printf("Error occurred while saving station snapshot: %s\n", err)
            errorCount++
        }
	}
    if (errorCount > 0) {
        fmt.Printf("%s: %d errors occurred\n\n", now.Format("Jan 2 2006 15:04:05"), errorCount)
    }
    fmt.Printf("%s: snapshots saved successfully\n", now.Format("Jan 2 2006 15:04:05"))
}

func (s *Service) DataCollectionWorker() {
	tickerDuration := time.Duration(s.config.DataCollectionIntervalMs) * time.Millisecond
	ticker := time.NewTicker(tickerDuration)
	s.saveStationSnapshots()
	for {
		select {
		case <-ticker.C:
			s.saveStationSnapshots()
		}
	}
}

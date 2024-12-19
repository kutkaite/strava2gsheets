package gsheets

import (
	"fmt"
	"math"
	"time"

	"strava2gsheets/internal/strava"
)

const (
	kmDistancePrecision = 2
)

type Stats struct {
	TrainingDateStr string
	TotalDur        string
	TotalDistance   float64
}

func Format(stats []strava.ActivityAgg) []Stats {
	formattedStats := make([]Stats, len(stats))
	i := 0
	for _, agg := range stats {
		duration := time.Duration(agg.Duration) * time.Second
		formattedStats[i] = Stats{
			TrainingDateStr: agg.TrainingDate.Format(layout),
			TotalDur:        formatDuration(duration),
			TotalDistance:   roundToKm(agg.Distance, kmDistancePrecision),
		}
		i++
	}
	return formattedStats
}

func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	// Return the formatted string in "HH.MM.SS" format
	return fmt.Sprintf("%02d.%02d.%02d", hours, minutes, seconds)
}

func roundToKm(val float64, precision uint) float64 {
	distanceInKm := val / 1000
	ratio := math.Pow(10, float64(precision))
	return math.Round(distanceInKm*ratio) / ratio
}

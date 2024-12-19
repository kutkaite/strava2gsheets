package strava

import (
	"sort"
	"time"
)

const layout = "2006-01-02"

type ActivityAgg struct {
	TrainingDate time.Time
	Duration     int64
	Distance     float64
}

func AggregatePerDay(activities []ActivitiesResponse) []ActivityAgg {
	aggDay := make(map[string]ActivityAgg)
	for _, activity := range activities {
		day := activity.StartDateLocal.Format(layout)
		if val, ok := aggDay[day]; ok {
			aggDay[day] = ActivityAgg{
				TrainingDate: activity.StartDateLocal,
				Duration:     val.Duration + int64(activity.ElapsedTime),
				Distance:     val.Distance + activity.Distance,
			}
		} else {
			aggDay[day] = ActivityAgg{
				TrainingDate: activity.StartDateLocal,
				Duration:     int64(activity.ElapsedTime),
				Distance:     activity.Distance,
			}
		}
	}

	stats := make([]ActivityAgg, len(aggDay))
	idx := 0
	for _, agg := range aggDay {
		stats[idx] = agg
		idx++
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].TrainingDate.Before(stats[j].TrainingDate)
	})
	return stats
}

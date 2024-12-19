package strava

import (
	"time"
)

type ActivitiesResponse struct {
	ResourceState int `json:"resource_state"`
	Athlete       struct {
		ID            int `json:"id"`
		ResourceState int `json:"resource_state"`
	} `json:"athlete"`
	Name               string      `json:"name"`
	Distance           float64     `json:"distance"`
	MovingTime         int         `json:"moving_time"`
	ElapsedTime        int         `json:"elapsed_time"`
	TotalElevationGain float64     `json:"total_elevation_gain"`
	Type               string      `json:"type"`
	SportType          string      `json:"sport_type"`
	ID                 int64       `json:"id"`
	StartDate          time.Time   `json:"start_date"`
	StartDateLocal     time.Time   `json:"start_date_local"`
	Timezone           string      `json:"timezone"`
	UtcOffset          float64     `json:"utc_offset"`
	LocationCity       interface{} `json:"location_city"`
	LocationState      interface{} `json:"location_state"`
	LocationCountry    string      `json:"location_country"`
	AchievementCount   int         `json:"achievement_count"`
	KudosCount         int         `json:"kudos_count"`
	CommentCount       int         `json:"comment_count"`
	AthleteCount       int         `json:"athlete_count"`
	PhotoCount         int         `json:"photo_count"`
	Map                struct {
		ID              string `json:"id"`
		SummaryPolyline string `json:"summary_polyline"`
		ResourceState   int    `json:"resource_state"`
	} `json:"map"`
	Trainer                    bool      `json:"trainer"`
	Commute                    bool      `json:"commute"`
	Manual                     bool      `json:"manual"`
	Private                    bool      `json:"private"`
	Visibility                 string    `json:"visibility"`
	Flagged                    bool      `json:"flagged"`
	GearID                     *string   `json:"gear_id"`
	StartLatlng                []float64 `json:"start_latlng"`
	EndLatlng                  []float64 `json:"end_latlng"`
	AverageSpeed               float64   `json:"average_speed"`
	MaxSpeed                   float64   `json:"max_speed"`
	HasHeartrate               bool      `json:"has_heartrate"`
	AverageHeartrate           float64   `json:"average_heartrate"`
	MaxHeartrate               float64   `json:"max_heartrate"`
	HeartrateOptOut            bool      `json:"heartrate_opt_out"`
	DisplayHideHeartrateOption bool      `json:"display_hide_heartrate_option"`
	ElevHigh                   float64   `json:"elev_high"`
	ElevLow                    float64   `json:"elev_low"`
	UploadID                   int64     `json:"upload_id"`
	UploadIDStr                string    `json:"upload_id_str"`
	ExternalID                 string    `json:"external_id"`
	FromAcceptedTag            bool      `json:"from_accepted_tag"`
	PrCount                    int       `json:"pr_count"`
	TotalPhotoCount            int       `json:"total_photo_count"`
	HasKudoed                  bool      `json:"has_kudoed"`
	SufferScore                float64   `json:"suffer_score"`
	WorkoutType                int       `json:"workout_type,omitempty"`
	AverageCadence             float64   `json:"average_cadence,omitempty"`
	AverageWatts               float64   `json:"average_watts,omitempty"`
	MaxWatts                   int       `json:"max_watts,omitempty"`
	WeightedAverageWatts       int       `json:"weighted_average_watts,omitempty"`
	DeviceWatts                bool      `json:"device_watts,omitempty"`
	Kilojoules                 float64   `json:"kilojoules,omitempty"`
}

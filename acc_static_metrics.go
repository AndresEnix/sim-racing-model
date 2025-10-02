package model

import (
	"reflect"
	"time"
)

// AccStaticMetrics is the name of the AccStatic metrics.
type AccStaticMetrics struct {
	ID                  uint      `gorm:"primaryKey;column:id"         json:"id"`
	UserID              string    `gorm:"column:user_id"               json:"userId"`
	SessionID           string    `gorm:"column:session_id"            json:"sessionId"`
	Timestamp           time.Time `gorm:"column:timestamp"             json:"timestamp"`
	SmVersion           string    `gorm:"column:sm_version"            json:"smVersion"`
	AcVersion           string    `gorm:"column:ac_version"            json:"acVersion"`
	NumberOfSessions    int64     `gorm:"column:number_of_sessions"    json:"numberOfSessions"`
	NumCars             int64     `gorm:"column:num_cars"              json:"numCars"`
	CarModel            string    `gorm:"column:car_model"             json:"carModel"`
	Track               string    `gorm:"column:track"                 json:"track"`
	PlayerName          string    `gorm:"column:player_name"           json:"playerName"`
	PlayerSurname       string    `gorm:"column:player_surname"        json:"playerSurname"`
	PlayerNick          string    `gorm:"column:player_nick"           json:"playerNick"`
	SectorCount         int64     `gorm:"column:sector_count"          json:"sectorCount"`
	MaxRpm              int64     `gorm:"column:max_rpm"               json:"maxRpm"`
	MaxFuel             float64   `gorm:"column:max_fuel"              json:"maxFuel"`
	PenaltiesEnabled    int64     `gorm:"column:penalties_enabled"     json:"penaltiesEnabled"`
	AidFuelRate         float64   `gorm:"column:aid_fuel_rate"         json:"aidFuelRate"`
	AidTireRate         float64   `gorm:"column:aid_tire_rate"         json:"aidTireRate"`
	AidMechanicalDamage float64   `gorm:"column:aid_mechanical_damage" json:"aidMechanicalDamage"`
	AllowTyreBlankets   float64   `gorm:"column:allow_tyre_blankets"   json:"allowTyreBlankets"`
	AidStability        float64   `gorm:"column:aid_stability"         json:"aidStability"`
	AidAutoClutch       int64     `gorm:"column:aid_auto_clutch"       json:"aidAutoClutch"`
	AidAutoBlip         int64     `gorm:"column:aid_auto_blip"         json:"aidAutoBlip"`
	PitWindowStart      int64     `gorm:"column:pit_window_start"      json:"pitWindowStart"`
	PitWindowEnd        int64     `gorm:"column:pit_window_end"        json:"pitWindowEnd"`
	IsOnline            int64     `gorm:"column:is_online"             json:"isOnline"`
	DryTyresName        string    `gorm:"column:dry_tyres_name"        json:"dryTyresName"`
	WetTyresName        string    `gorm:"column:wet_tyres_name"        json:"wetTyresName"`
}

// NewAccStaticMetrics creates a new AccStaticMetrics struct with default values.
func NewAccStaticMetrics() *AccStaticMetrics {
	return &AccStaticMetrics{
		ID:                  0,
		UserID:              "",
		SessionID:           "",
		Timestamp:           time.Time{},
		SmVersion:           "",
		AcVersion:           "",
		NumberOfSessions:    0,
		NumCars:             0,
		CarModel:            "",
		Track:               "",
		PlayerName:          "",
		PlayerSurname:       "",
		PlayerNick:          "",
		SectorCount:         0,
		MaxRpm:              0,
		MaxFuel:             0.0,
		PenaltiesEnabled:    0,
		AidFuelRate:         0.0,
		AidTireRate:         0.0,
		AidMechanicalDamage: 0.0,
		AllowTyreBlankets:   0.0,
		AidStability:        0.0,
		AidAutoClutch:       0,
		AidAutoBlip:         0,
		PitWindowStart:      0,
		PitWindowEnd:        0,
		IsOnline:            0,
		DryTyresName:        "",
		WetTyresName:        "",
	}
}

// New creates a new AccStaticMetrics struct with default values.
func (metric *AccStaticMetrics) New() Metrics {
	return NewAccStaticMetrics()
}

// Game returns the name of the game for the AccStaticMetrics metric.
func (metric *AccStaticMetrics) Game() string {
	return AssettoCorsaCompetizione
}

// Name returns the name of the AccStaticMetrics metric.
func (metric *AccStaticMetrics) Name() string {
	return AccStaticMetricsName
}

// DataPoints returns the field names of the AccStaticMetrics struct.
func (metric *AccStaticMetrics) DataPoints() []string {
	metricType := reflect.TypeOf(metric)

	if metricType.Kind() == reflect.Ptr {
		metricType = metricType.Elem()
	}

	if metricType.Kind() != reflect.Struct {
		return []string{}
	}

	var names []string

	for i := range metricType.NumField() {
		field := metricType.Field(i)

		if field.IsExported() {
			names = append(names, field.Name)
		}
	}

	return names
}

// TableName returns the name of the DB table for the AccStaticMetrics metric.
func (metric *AccStaticMetrics) TableName() string {
	return "acc_static_metrics"
}

// AddSessionInfo adds the user ID and session ID to the metric.
func (metric *AccStaticMetrics) AddSessionInfo(userID, sessionID string) {
	metric.UserID = userID
	metric.SessionID = sessionID
}

package model


import (
	"reflect"
	"time"
)


// AccStaticMetrics is the name of the AccStatic metrics.
type AccStaticMetrics struct {
	ID                  string    `json:"id"`
	UserID              string    `json:"userId"`
	SessionID           string    `json:"sessionId"`
	Timestamp           time.Time `json:"timestamp"`
	SmVersion           string    `json:"smVersion"`
	AcVersion           string    `json:"acVersion"`
	NumberOfSessions    int64     `json:"numberOfSessions"`
	NumCars             int64     `json:"numCars"`
	CarModel            string    `json:"carModel"`
	Track               string    `json:"track"`
	PlayerName          string    `json:"playerName"`
	PlayerSurname       string    `json:"playerSurname"`
	PlayerNick          string    `json:"playerNick"`
	SectorCount         int64     `json:"sectorCount"`
	MaxRpm              int64     `json:"maxRpm"`
	MaxFuel             float64   `json:"maxFuel"`
	PenaltiesEnabled    int64     `json:"penaltiesEnabled"`
	AidFuelRate         float64   `json:"aidFuelRate"`
	AidTireRate         float64   `json:"aidTireRate"`
	AidMechanicalDamage float64   `json:"aidMechanicalDamage"`
	AllowTyreBlankets   float64   `json:"allowTyreBlankets"`
	AidStability        float64   `json:"aidStability"`
	AidAutoClutch       int64     `json:"aidAutoClutch"`
	AidAutoBlip         int64     `json:"aidAutoBlip"`
	PitWindowStart      int64     `json:"pitWindowStart"`
	PitWindowEnd        int64     `json:"pitWindowEnd"`
	IsOnline            int64     `json:"isOnline"`
	DryTyresName        string    `json:"dryTyresName"`
	WetTyresName        string    `json:"wetTyresName"`
}


// NewAccStaticMetrics creates a new AccStaticMetrics struct with default values.
func NewAccStaticMetrics() *AccStaticMetrics {
	return &AccStaticMetrics{
		ID:                  "",
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
	v := reflect.ValueOf(metric)
	t := v.Type()

	fieldNames := make([]string, 0, v.NumField())
	for i := range v.NumField() {
		fieldNames = append(fieldNames, t.Field(i).Name)
	}

	return fieldNames
}


// AddSessionInfo adds the user ID and session ID to the metric.
func (metric *AccStaticMetrics) AddSessionInfo(userID, sessionID string) {
	metric.UserID = userID
	metric.SessionID = sessionID
}

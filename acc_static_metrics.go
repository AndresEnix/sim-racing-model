package model

import "time"

type AccStaticMetrics struct {
	Id                  string    `json:"Id"`
	UserId              string    `json:"UserId"`
	SessionId           string    `json:"SessionId"`
	Timestamp           time.Time `json:"Timestamp"`
	SmVersion           string    `json:"SmVersion"`
	AcVersion           string    `json:"AcVersion"`
	NumberOfSessions    int64     `json:"NumberOfSessions"`
	NumCars             int64     `json:"NumCars"`
	CarModel            string    `json:"CarModel"`
	Track               string    `json:"Track"`
	PlayerName          string    `json:"PlayerName"`
	PlayerSurname       string    `json:"PlayerSurname"`
	PlayerNick          string    `json:"PlayerNick"`
	SectorCount         int64     `json:"SectorCount"`
	MaxRpm              int64     `json:"MaxRpm"`
	MaxFuel             float64   `json:"MaxFuel"`
	PenaltiesEnabled    int64     `json:"PenaltiesEnabled"`
	AidFuelRate         float64   `json:"AidFuelRate"`
	AidTireRate         float64   `json:"AidTireRate"`
	AidMechanicalDamage float64   `json:"AidMechanicalDamage"`
	AllowTyreBlankets   float64   `json:"AllowTyreBlankets"`
	AidStability        float64   `json:"AidStability"`
	AidAutoClutch       int64     `json:"AidAutoClutch"`
	AidAutoBlip         int64     `json:"AidAutoBlip"`
	PitWindowStart      int64     `json:"PitWindowStart"`
	PitWindowEnd        int64     `json:"PitWindowEnd"`
	IsOnline            int64     `json:"IsOnline"`
	DryTyresName        string    `json:"DryTyresName"`
	WetTyresName        string    `json:"WetTyresName"`
}

func (metric *AccStaticMetrics) Game() string {
	return ASSETTO_CORSA_COMPETIZIONE
}

func (metric *AccStaticMetrics) Name() string {
	return ACC_STATIC_METRICS_NAME
}

func (metric *AccStaticMetrics) DataPoints() []string {
	return getStructFieldNames(*metric)
}

func (metric *AccStaticMetrics) AddSessionInfo(userId, sessionId string) {
	metric.UserId = userId
	metric.SessionId = sessionId
}

func (metric *AccStaticMetrics) New() Metrics {
	return &AccStaticMetrics{}
}

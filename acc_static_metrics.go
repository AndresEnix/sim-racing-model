package model

import "time"

type AccStaticMetrics struct {
	Id                  string    `json:"Id"`
	UserId              string    `json:"UserId"`
	SessionId           string    `json:"SessionId"`
	Timestamp           time.Time `json:"Timestamp"`
	SmVersion           string    `json:"SmVersion"`
	AcVersion           string    `json:"AcVersion"`
	NumberOfSessions    int32     `json:"NumberOfSessions"`
	NumCars             int32     `json:"NumCars"`
	CarModel            string    `json:"CarModel"`
	Track               string    `json:"Track"`
	PlayerName          string    `json:"PlayerName"`
	PlayerSurname       string    `json:"PlayerSurname"`
	PlayerNick          string    `json:"PlayerNick"`
	SectorCount         int32     `json:"SectorCount"`
	MaxRpm              int32     `json:"MaxRpm"`
	MaxFuel             float32   `json:"MaxFuel"`
	PenaltiesEnabled    int32     `json:"PenaltiesEnabled"`
	AidFuelRate         float32   `json:"AidFuelRate"`
	AidTireRate         float32   `json:"AidTireRate"`
	AidMechanicalDamage float32   `json:"AidMechanicalDamage"`
	AllowTyreBlankets   float32   `json:"AllowTyreBlankets"`
	AidStability        float32   `json:"AidStability"`
	AidAutoClutch       int32     `json:"AidAutoClutch"`
	AidAutoBlip         int32     `json:"AidAutoBlip"`
	PitWindowStart      int32     `json:"PitWindowStart"`
	PitWindowEnd        int32     `json:"PitWindowEnd"`
	IsOnline            int32     `json:"IsOnline"`
	DryTyresName        string    `json:"DryTyresName"`
	WetTyresName        string    `json:"WetTyresName"`
}

func (metric AccStaticMetrics) Game() string {
	return ASSETTO_CORSA_COMPETIZIONE
}

func (metric AccStaticMetrics) Name() string {
	return ACC_STATIC_METRICS_NAME
}

func (metric AccStaticMetrics) DataPoints() []string {
	return getStructFieldNames(metric)
}

func (metric AccStaticMetrics) SetFingerprint(userId, sessionId string, timestamp time.Time) {
	metric.UserId = userId
	metric.SessionId = sessionId
	metric.Timestamp = timestamp
}

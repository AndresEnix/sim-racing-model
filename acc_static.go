package model

import (
	"fmt"
	"unsafe"
)

type AccStaticMemory struct {
	SmVersion                [15]uint16
	AcVersion                [15]uint16
	NumberOfSessions         int32
	NumCars                  int32
	CarModel                 [33]uint16
	Track                    [33]uint16
	PlayerName               [33]uint16
	PlayerSurname            [33]uint16
	PlayerNick               [33]uint16
	SectorCount              int32
	MaxTorque                float32 // Not used
	MaxPower                 float32 // Not used
	MaxRpm                   int32
	MaxFuel                  float32
	SuspensionMaxTravel      [4]float32 // Not used
	TyreRadius               [4]float32 // Not used
	MaxTurboBoost            float32    // Not used
	Deprecated_1             float32    // Not used
	Deprecated_2             float32    // Not used
	PenaltiesEnabled         int32
	AidFuelRate              float32
	AidTireRate              float32
	AidMechanicalDamage      float32
	AllowTyreBlankets        float32
	AidStability             float32
	AidAutoclutch            int32
	AidAutoBlip              int32
	HasDRS                   int32      // Not used
	HasERS                   int32      // Not used
	HasKERS                  int32      // Not used
	KersMaxJ                 float32    // Not used
	EngineBrakeSettingsCount int32      // Not used
	ErsPowerControllerCount  int32      // Not used
	TrackSplineLength        float32    // Not used
	TrackConfiguration       [33]uint16 // Not used
	ErsMaxJ                  float32    // Not used
	IsTimedRace              int32      // Not used
	HasExtraLap              int32      // Not used
	CarSkin                  [33]uint16 // Not used
	ReversedGridPositions    int32      // Not used
	PitWindowStart           int32
	PitWindowEnd             int32
	IsOnline                 int32
	DryTyresName             [33]uint16
	WetTyresName             [33]uint16
}

func (memory AccStaticMemory) GetFileName() string {
	return "acpmf_static"
}

func (memory AccStaticMemory) GetFilePath() string {
	return fmt.Sprintf("Local\\%s", memory.GetFileName())
}

func (memory AccStaticMemory) Create(pointer uintptr) DataMemoryMapping {
	return (*AccStaticMemory)(unsafe.Pointer(pointer))
}

func (memory AccStaticMemory) ToMetric() Metric {
	return AccStaticMetric{
		SmVersion:           uint16ToString(memory.SmVersion[:]),
		AcVersion:           uint16ToString(memory.AcVersion[:]),
		NumberOfSessions:    memory.NumberOfSessions,
		NumCars:             memory.NumCars,
		CarModel:            uint16ToString(memory.CarModel[:]),
		Track:               uint16ToString(memory.Track[:]),
		PlayerName:          uint16ToString(memory.PlayerName[:]),
		PlayerSurname:       uint16ToString(memory.PlayerSurname[:]),
		PlayerNick:          uint16ToString(memory.PlayerNick[:]),
		SectorCount:         memory.SectorCount,
		MaxRpm:              memory.MaxRpm,
		MaxFuel:             memory.MaxFuel,
		PenaltiesEnabled:    memory.PenaltiesEnabled,
		AidFuelRate:         memory.AidFuelRate,
		AidTireRate:         memory.AidTireRate,
		AidMechanicalDamage: memory.AidMechanicalDamage,
		AllowTyreBlankets:   memory.AllowTyreBlankets,
		AidStability:        memory.AidStability,
		AidAutoclutch:       memory.AidAutoclutch,
		AidAutoBlip:         memory.AidAutoBlip,
		PitWindowStart:      memory.PitWindowStart,
		PitWindowEnd:        memory.PitWindowEnd,
		IsOnline:            memory.IsOnline,
		DryTyresName:        uint16ToString(memory.DryTyresName[:]),
		WetTyresName:        uint16ToString(memory.WetTyresName[:]),
	}
}

type AccStaticMetric struct {
	SmVersion           string  `json:"SmVersion"`
	AcVersion           string  `json:"AcVersion"`
	NumberOfSessions    int32   `json:"NumberOfSessions"`
	NumCars             int32   `json:"NumCars"`
	CarModel            string  `json:"CarModel"`
	Track               string  `json:"Track"`
	PlayerName          string  `json:"PlayerName"`
	PlayerSurname       string  `json:"PlayerSurname"`
	PlayerNick          string  `json:"PlayerNick"`
	SectorCount         int32   `json:"SectorCount"`
	MaxRpm              int32   `json:"MaxRpm"`
	MaxFuel             float32 `json:"MaxFuel"`
	PenaltiesEnabled    int32   `json:"PenaltiesEnabled"`
	AidFuelRate         float32 `json:"AidFuelRate"`
	AidTireRate         float32 `json:"AidTireRate"`
	AidMechanicalDamage float32 `json:"AidMechanicalDamage"`
	AllowTyreBlankets   float32 `json:"AllowTyreBlankets"`
	AidStability        float32 `json:"AidStability"`
	AidAutoclutch       int32   `json:"AidAutoclutch"`
	AidAutoBlip         int32   `json:"AidAutoBlip"`
	PitWindowStart      int32   `json:"PitWindowStart"`
	PitWindowEnd        int32   `json:"PitWindowEnd"`
	IsOnline            int32   `json:"IsOnline"`
	DryTyresName        string  `json:"DryTyresName"`
	WetTyresName        string  `json:"WetTyresName"`
}

func (metric AccStaticMetric) Print() {
	fmt.Println("")
}

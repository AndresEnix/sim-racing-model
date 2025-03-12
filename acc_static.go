package model

import (
	"fmt"
	"unsafe"
	"golang.org/x/sys/windows"
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

func (memory AccStaticMemory) FileName() string {
	return "Local\\acpmf_static"
}

func (memory AccStaticMemory) Create(pointer uintptr) DataMemoryMapping {
	return (*AccStaticMemory)(unsafe.Pointer(pointer))
}

func (memory AccStaticMemory) ToMetric() Metric {
	return AccStaticMetric{
		SmVersion:           windows.UTF16ToString(memory.SmVersion[:]),
		AcVersion:           windows.UTF16ToString(memory.AcVersion[:]),
		NumberOfSessions:    memory.NumberOfSessions,
		NumCars:             memory.NumCars,
		CarModel:            windows.UTF16ToString(memory.CarModel[:]),
		Track:               windows.UTF16ToString(memory.Track[:]),
		PlayerName:          windows.UTF16ToString(memory.PlayerName[:]),
		PlayerSurname:       windows.UTF16ToString(memory.PlayerSurname[:]),
		PlayerNick:          windows.UTF16ToString(memory.PlayerNick[:]),
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
		DryTyresName:        windows.UTF16ToString(memory.DryTyresName[:]),
		WetTyresName:        windows.UTF16ToString(memory.WetTyresName[:]),
	}
}

type AccStaticMetric struct {
	SmVersion           string
	AcVersion           string
	NumberOfSessions    int32
	NumCars             int32
	CarModel            string
	Track               string
	PlayerName          string
	PlayerSurname       string
	PlayerNick          string
	SectorCount         int32
	MaxRpm              int32
	MaxFuel             float32
	PenaltiesEnabled    int32
	AidFuelRate         float32
	AidTireRate         float32
	AidMechanicalDamage float32
	AllowTyreBlankets   float32
	AidStability        float32
	AidAutoclutch       int32
	AidAutoBlip         int32
	PitWindowStart      int32
	PitWindowEnd        int32
	IsOnline            int32
	DryTyresName        string
	WetTyresName        string
}

func (metric AccStaticMetric) Print() {
	fmt.Println("")
}

package model

import (
	"fmt"
	"hash/fnv"
	"time"
	"unsafe"
)

type AccStaticData struct {
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
	AidAutoClutch            int32
	AidAutoBlip              int32
	HasDRS                   int32      // Not used
	HasERS                   int32      // Not used
	HasKers                  int32      // Not used
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

func (memory *AccStaticData) Name() string {
	return ACC_STATIC_FILE_NAME
}

func (memory *AccStaticData) Path() string {
	return ACC_FILES_PREFIX + memory.Name()
}

func (memory *AccStaticData) ReadFrequency() time.Duration {
	return time.Duration(getUint64Env(memory.Name()+READ_FREQUENCY_SUFFIX, ACC_STATIC_DEFAULT_READ_FREQUENCY)) * time.Millisecond
}

func (memory *AccStaticData) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(uint16ToString(memory.SmVersion[:])))
	h.Write([]byte(uint16ToString(memory.AcVersion[:])))
	h.Write([]byte(fmt.Sprintf("%d", memory.NumberOfSessions)))
	h.Write([]byte(fmt.Sprintf("%d", memory.NumCars)))
	h.Write([]byte(uint16ToString(memory.CarModel[:])))
	h.Write([]byte(uint16ToString(memory.Track[:])))
	h.Write([]byte(uint16ToString(memory.PlayerName[:])))
	h.Write([]byte(uint16ToString(memory.PlayerSurname[:])))
	h.Write([]byte(uint16ToString(memory.PlayerNick[:])))
	return h.Sum32()
}

func (memory *AccStaticData) MapValues(pointer uintptr) {
	newValue := (*AccStaticData)(unsafe.Pointer(pointer))
	*memory = *newValue
}

func (memory *AccStaticData) CreateMetric() Metrics {
	return &AccStaticMetrics{
		Timestamp:           time.Now().UTC(),
		SmVersion:           uint16ToString(memory.SmVersion[:]),
		AcVersion:           uint16ToString(memory.AcVersion[:]),
		NumberOfSessions:    int64(memory.NumberOfSessions),
		NumCars:             int64(memory.NumCars),
		CarModel:            uint16ToString(memory.CarModel[:]),
		Track:               uint16ToString(memory.Track[:]),
		PlayerName:          uint16ToString(memory.PlayerName[:]),
		PlayerSurname:       uint16ToString(memory.PlayerSurname[:]),
		PlayerNick:          uint16ToString(memory.PlayerNick[:]),
		SectorCount:         int64(memory.SectorCount),
		MaxRpm:              int64(memory.MaxRpm),
		MaxFuel:             float64(memory.MaxFuel),
		PenaltiesEnabled:    int64(memory.PenaltiesEnabled),
		AidFuelRate:         float64(memory.AidFuelRate),
		AidTireRate:         float64(memory.AidTireRate),
		AidMechanicalDamage: float64(memory.AidMechanicalDamage),
		AllowTyreBlankets:   float64(memory.AllowTyreBlankets),
		AidStability:        float64(memory.AidStability),
		AidAutoClutch:       int64(memory.AidAutoClutch),
		AidAutoBlip:         int64(memory.AidAutoBlip),
		PitWindowStart:      int64(memory.PitWindowStart),
		PitWindowEnd:        int64(memory.PitWindowEnd),
		IsOnline:            int64(memory.IsOnline),
		DryTyresName:        uint16ToString(memory.DryTyresName[:]),
		WetTyresName:        uint16ToString(memory.WetTyresName[:]),
	}
}

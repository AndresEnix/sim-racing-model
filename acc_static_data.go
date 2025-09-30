package model


import (
	"fmt"
	"hash/fnv"
	"log"
	"math"
	"time"
	"unsafe"
	"encoding/json"
)


// AccStaticData represents the structure of the ACC graphics data shared memory.
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
	MaxTorque                float32 // Not used.
	MaxPower                 float32 // Not used.
	MaxRpm                   int32
	MaxFuel                  float32
	SuspensionMaxTravel      [4]float32 // Not used.
	TyreRadius               [4]float32 // Not used.
	MaxTurboBoost            float32    // Not used.
	Deprecated1              float32    // Not used.
	Deprecated2              float32    // Not used.
	PenaltiesEnabled         int32
	AidFuelRate              float32
	AidTireRate              float32
	AidMechanicalDamage      float32
	AllowTyreBlankets        float32
	AidStability             float32
	AidAutoClutch            int32
	AidAutoBlip              int32
	HasDRS                   int32      // Not used.
	HasERS                   int32      // Not used.
	HasKers                  int32      // Not used.
	KersMaxJ                 float32    // Not used.
	EngineBrakeSettingsCount int32      // Not used.
	ErsPowerControllerCount  int32      // Not used.
	TrackSplineLength        float32    // Not used.
	TrackConfiguration       [33]uint16 // Not used.
	ErsMaxJ                  float32    // Not used.
	IsTimedRace              int32      // Not used.
	HasExtraLap              int32      // Not used.
	CarSkin                  [33]uint16 // Not used.
	ReversedGridPositions    int32      // Not used.
	PitWindowStart           int32
	PitWindowEnd             int32
	IsOnline                 int32
	DryTyresName             [33]uint16
	WetTyresName             [33]uint16
}


// NewAccStaticData creates a new instance of AccStaticData with default values.
func NewAccStaticData() *AccStaticData {
	return &AccStaticData{
		SmVersion:                [15]uint16{},
		AcVersion:                [15]uint16{},
		NumberOfSessions:         0,
		NumCars:                  0,
		CarModel:                 [33]uint16{},
		Track:                    [33]uint16{},
		PlayerName:               [33]uint16{},
		PlayerSurname:            [33]uint16{},
		PlayerNick:               [33]uint16{},
		SectorCount:              0,
		MaxTorque:                0.0,
		MaxPower:                 0.0,
		MaxRpm:                   0,
		MaxFuel:                  0.0,
		SuspensionMaxTravel:      [4]float32{},
		TyreRadius:               [4]float32{},
		MaxTurboBoost:            0.0,
		Deprecated1:              0.0,
		Deprecated2:              0.0,
		PenaltiesEnabled:         0,
		AidFuelRate:              0.0,
		AidTireRate:              0.0,
		AidMechanicalDamage:      0.0,
		AllowTyreBlankets:        0.0,
		AidStability:             0.0,
		AidAutoClutch:            0,
		AidAutoBlip:              0,
		HasDRS:                   0,
		HasERS:                   0,
		HasKers:                  0,
		KersMaxJ:                 0.0,
		EngineBrakeSettingsCount: 0,
		ErsPowerControllerCount:  0,
		TrackSplineLength:        0.0,
		TrackConfiguration:       [33]uint16{},
		ErsMaxJ:                  0.0,
		IsTimedRace:              0,
		HasExtraLap:              0,
		CarSkin:                  [33]uint16{},
		ReversedGridPositions:    0,
		PitWindowStart:           0,
		PitWindowEnd:             0,
		IsOnline:                 0,
		DryTyresName:             [33]uint16{},
		WetTyresName:             [33]uint16{},
	}
}


// Name returns the name of the shared memory file for ACC static data.
func (memory *AccStaticData) Name() string {
	return AccStaticFileName
}


// Path returns the full path to the shared memory file for ACC static data.
func (memory *AccStaticData) Path() string {
	return AccFilesPrefix + memory.Name()
}


// ReadFrequency returns the frequency at which the ACC static data should be read.
func (memory *AccStaticData) ReadFrequency() time.Duration {
	frequencyMs := getUint64Env(memory.Name()+ReadFrequencySuffix, AccStaticDefaultReadFrequency)
	if frequencyMs > uint64(MaxAllowedMs) {
		log.Printf(
			"WARNING: Duration value (%d ms) from %s overflows time.Duration. Capping at MaxDuration.\n",
			frequencyMs,
			memory.Name()+ReadFrequencySuffix,
		)

		return time.Duration(math.MaxInt64)
	}

	return time.Duration(frequencyMs) * time.Millisecond
}


// Hash generates a hash based on a combination of fields of the ACC static data.
func (memory *AccStaticData) Hash() uint32 {
	hasher := fnv.New32a()
	
	_, err := hasher.Write([]byte(uint16ToString(memory.SmVersion[:])))
	if err != nil {
		log.Printf("WARNING: Failed to use SmVersion in %s hash: %s\n", memory.Name(), err.Error())
	}

	_, err = hasher.Write([]byte(uint16ToString(memory.AcVersion[:])))
	if err != nil {
		log.Printf("WARNING: Failed to use AcVersion in %s hash: %s\n", memory.Name(), err.Error())
	}

	_, err = fmt.Fprintf(hasher, "%d", memory.NumberOfSessions)
	if err != nil {
		log.Printf("WARNING: Failed to use NumberOfSessions in %s hash: %s\n", memory.Name(), err.Error())
	}

	_, err = fmt.Fprintf(hasher, "%d", memory.NumCars)
	if err != nil {
		log.Printf("WARNING: Failed to use NumCars in %s hash: %s\n", memory.Name(), err.Error())
	}

	_, err = hasher.Write([]byte(uint16ToString(memory.CarModel[:])))
	if err != nil {
		log.Printf("WARNING: Failed to use CarModel in %s hash: %s\n", memory.Name(), err.Error())
	}

	_, err = hasher.Write([]byte(uint16ToString(memory.Track[:])))
	if err != nil {
		log.Printf("WARNING: Failed to use Track in %s hash: %s\n", memory.Name(), err.Error())
	}

	_, err = hasher.Write([]byte(uint16ToString(memory.PlayerName[:])))
	if err != nil {
		log.Printf("WARNING: Failed to use PlayerName in %s hash: %s\n", memory.Name(), err.Error())
	}

	_, err = hasher.Write([]byte(uint16ToString(memory.PlayerSurname[:])))
	if err != nil {
		log.Printf("WARNING: Failed to use PlayerSurname in %s hash: %s\n", memory.Name(), err.Error())
	}

	_, err = hasher.Write([]byte(uint16ToString(memory.PlayerNick[:])))
	if err != nil {
		log.Printf("WARNING: Failed to use PlayerNick in %s hash: %s\n", memory.Name(), err.Error())
	}

	return hasher.Sum32()
}


//nolint:gosec
// MapValues maps the values from a memory pointer to the AccStaticData struct.
func (memory *AccStaticData) MapValues(pointer uintptr) {
	newValue := (*AccStaticData)(unsafe.Pointer(pointer))//nolint:govet
	*memory = *newValue
}


// CreateMetricsJSON creates JSON bytes from the current AccStaticData.
func (memory *AccStaticData) CreateMetricsJSON() ([]byte, error) {
	accStaticMetrics := &AccStaticMetrics{
		ID:                  "",
		UserID:              "",
		SessionID:           "",
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
	
	metrics, err := json.Marshal(accStaticMetrics)
	if err != nil {
        return nil, fmt.Errorf("failed to marshal AccStaticData to JSON: %w", err)
    }
    
    return metrics, nil
}

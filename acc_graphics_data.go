package model


import (
	"fmt"
	"hash/fnv"
	"log"
	"math"
	"time"
	"unsafe"
)


// AccGraphicsData represents the structure of the ACC graphics data shared memory.
type AccGraphicsData struct {
	PacketID                 int32
	Status                   int32
	Session                  int32
	CurrentTime              [15]uint16
	LastTime                 [15]uint16
	BestTime                 [15]uint16
	Split                    [15]uint16
	CompletedLaps            int32
	Position                 int32
	ICurrentTime             int32
	ILastTime                int32
	IBestTime                int32
	SessionTimeLeft          float32
	DistanceTraveled         float32
	IsInPit                  int32
	CurrentSectorIndex       int32
	LastSectorTime           int32
	NumberOfLaps             int32
	TyreCompound             [33]uint16
	ReplayTimeMultiplier     float32 // Not used.
	NormalizedCarPosition    float32
	ActiveCars               int32
	AarCoordinates           [60][3]float32
	CarID                    [60]int32
	PlayerCarID              int32
	PenaltyTime              float32
	Flag                     int32
	Penalty                  int32
	IdealLineOn              int32
	IsInPitLane              int32
	SurfaceGrip              float32
	MandatoryPitDone         int32
	WindSpeed                float32
	WindDirection            float32
	IsSetupMenuVisible       int32
	MainDisplayIndex         int32
	SecondaryDisplayIndex    int32
	Tc                       int32
	TcCut                    int32
	EngineMap                int32
	Abs                      int32
	FuelXLap                 float32
	RainLights               int32
	FlashingLights           int32
	LightsStage              int32
	ExhaustTemperature       float32
	WiperLV                  int32
	DriverStintTotalTimeLeft int32
	DriverStintTimeLeft      int32
	RainTyres                int32
	SessionIndex             int32
	UsedFuel                 float32
	DeltaLapTime             [15]uint16
	IDeltaLapTime            int32
	EstimatedLapTime         [15]uint16
	IEstimatedLapTime        int32
	IsDeltaPositive          int32
	ISplit                   int32
	IsValidLap               int32
	FuelEstimatedLaps        float32
	TrackStatus              [33]uint16
	MissingMandatoryPits     int32
	Clock                    float32
	DirectionLightsLeft      int32
	DirectionLightsRight     int32
	GlobalYellow             int32
	GlobalYellow1            int32
	GlobalYellow2            int32
	GlobalYellow3            int32
	GlobalWhite              int32
	GlobalGreen              int32
	GlobalChequered          int32
	GlobalRed                int32
	MfdTyreSet               int32
	MfdFuelToAdd             float32
	MfdTyrePressureLF        float32
	MfdTyrePressureRF        float32
	MfdTyrePressureLR        float32
	MfdTyrePressureRR        float32
	TrackGripStatus          int32
	RainIntensity            int32
	RainIntensityIn10min     int32
	RainIntensityIn30min     int32
	CurrentTyreSet           int32
	StrategyTyreSet          int32
	GapAhead                 int32
	GapBehind                int32
}


//nolint:funlen
// NewAccGraphicsData creates a new instance of AccGraphicsData with default values.
func NewAccGraphicsData() *AccGraphicsData {
	return &AccGraphicsData{
		PacketID:                 0,
		Status:                   0,
		Session:                  0,
		CurrentTime:              [15]uint16{},
		LastTime:                 [15]uint16{},
		BestTime:                 [15]uint16{},
		Split:                    [15]uint16{},
		CompletedLaps:            0,
		Position:                 0,
		ICurrentTime:             0,
		ILastTime:                0,
		IBestTime:                0,
		SessionTimeLeft:          0,
		DistanceTraveled:         0.0,
		IsInPit:                  0,
		CurrentSectorIndex:       0,
		LastSectorTime:           0,
		NumberOfLaps:             0,
		TyreCompound:             [33]uint16{},
		ReplayTimeMultiplier:     0.0,
		NormalizedCarPosition:    0.0,
		ActiveCars:               0,
		AarCoordinates:           [60][3]float32{},
		CarID:                    [60]int32{},
		PlayerCarID:              0,
		PenaltyTime:              0.0,
		Flag:                     0,
		Penalty:                  0,
		IdealLineOn:              0,
		IsInPitLane:              0,
		SurfaceGrip:              0.0,
		MandatoryPitDone:         0,
		WindSpeed:                0.0,
		WindDirection:            0.0,
		IsSetupMenuVisible:       0,
		MainDisplayIndex:         0,
		SecondaryDisplayIndex:    0,
		Tc:                       0,
		TcCut:                    0,
		EngineMap:                0,
		Abs:                      0,
		FuelXLap:                 0.0,
		RainLights:               0,
		FlashingLights:           0,
		LightsStage:              0,
		ExhaustTemperature:       0.0,
		WiperLV:                  0,
		DriverStintTotalTimeLeft: 0,
		DriverStintTimeLeft:      0,
		RainTyres:                0,
		SessionIndex:             0,
		UsedFuel:                 0.0,
		DeltaLapTime:             [15]uint16{},
		IDeltaLapTime:            0,
		EstimatedLapTime:         [15]uint16{},
		IEstimatedLapTime:        0,
		IsDeltaPositive:          0,
		ISplit:                   0,
		IsValidLap:               0,
		FuelEstimatedLaps:        0.0,
		TrackStatus:              [33]uint16{},
		MissingMandatoryPits:     0,
		Clock:                    0.0,
		DirectionLightsLeft:      0,
		DirectionLightsRight:     0,
		GlobalYellow:             0,
		GlobalYellow1:            0,
		GlobalYellow2:            0,
		GlobalYellow3:            0,
		GlobalWhite:              0,
		GlobalGreen:              0,
		GlobalChequered:          0,
		GlobalRed:                0,
		MfdTyreSet:               0,
		MfdFuelToAdd:             0.0,
		MfdTyrePressureLF:        0.0,
		MfdTyrePressureRF:        0.0,
		MfdTyrePressureLR:        0.0,
		MfdTyrePressureRR:        0.0,
		TrackGripStatus:          0,
		RainIntensity:            0,
		RainIntensityIn10min:     0,
		RainIntensityIn30min:     0,
		CurrentTyreSet:           0,
		StrategyTyreSet:          0,
		GapAhead:                 0,
		GapBehind:                0,
	}
}


// Name returns the name of the shared memory file for ACC graphics data.
func (memory *AccGraphicsData) Name() string {
	return AccGraphicsFileName
}


// Path returns the full path to the shared memory file for ACC graphics data.
func (memory *AccGraphicsData) Path() string {
	return AccFilesPrefix + memory.Name()
}


// ReadFrequency returns the frequency at which the ACC graphics data should be read.
func (memory *AccGraphicsData) ReadFrequency() time.Duration {
	frequencyMs := getUint64Env(memory.Name()+ReadFrequencySuffix, AccGraphicsDefaultReadFrequency)
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


// Hash generates a hash based on the PacketID of the ACC graphics data.
func (memory *AccGraphicsData) Hash() uint32 {
	hasher := fnv.New32a()

	_, err := fmt.Fprintf(hasher, "%d", memory.PacketID)
	if err != nil {
		log.Printf("WARNING: Failed to generate hash for %s: %s\n", memory.Name(), err.Error())

		return 0
	}

	return hasher.Sum32()
}


//nolint:gosec
// MapValues maps the values from a memory pointer to the AccGraphicsData struct.
func (memory *AccGraphicsData) MapValues(pointer uintptr) {
	newValue := (*AccGraphicsData)(unsafe.Pointer(pointer))//nolint:govet
	*memory = *newValue
}


//nolint:funlen
// CreateMetric creates a new AccGraphicsMetrics instance from the current AccGraphicsData.
func (memory *AccGraphicsData) CreateMetric() *AccGraphicsMetrics {
	return &AccGraphicsMetrics{
		ID:                       "",
		UserID:                   "",
		SessionID:                "",
		Timestamp:                time.Now().UTC(),
		PacketID:                 int64(memory.PacketID),
		Status:                   int64(memory.Status),
		Session:                  int64(memory.Session),
		CurrentTime:              uint16ToString(memory.CurrentTime[:]),
		LastTime:                 uint16ToString(memory.LastTime[:]),
		BestTime:                 uint16ToString(memory.BestTime[:]),
		Split:                    uint16ToString(memory.Split[:]),
		CompletedLaps:            int64(memory.CompletedLaps),
		Position:                 int64(memory.Position),
		ICurrentTime:             int64(memory.ICurrentTime),
		ILastTime:                int64(memory.ILastTime),
		IBestTime:                int64(memory.IBestTime),
		SessionTimeLeft:          float64(memory.SessionTimeLeft),
		DistanceTraveled:         float64(memory.DistanceTraveled),
		IsInPit:                  int64(memory.IsInPit),
		CurrentSectorIndex:       int64(memory.CurrentSectorIndex),
		LastSectorTime:           int64(memory.LastSectorTime),
		NumberOfLaps:             int64(memory.NumberOfLaps),
		TyreCompound:             uint16ToString(memory.TyreCompound[:]),
		NormalizedCarPosition:    float64(memory.NormalizedCarPosition),
		ActiveCars:               int64(memory.ActiveCars),
		AarCoordinates:           twoDimensionSliceFloat32To64(floatArray60_3ToSlice(memory.AarCoordinates)),
		CarID:                    oneDimensionSliceInt32To64(memory.CarID[:]),
		PlayerCarID:              int64(memory.PlayerCarID),
		PenaltyTime:              float64(memory.PenaltyTime),
		Flag:                     int64(memory.Flag),
		Penalty:                  int64(memory.Penalty),
		IdealLineOn:              int64(memory.IdealLineOn),
		IsInPitLane:              int64(memory.IsInPitLane),
		SurfaceGrip:              float64(memory.SurfaceGrip),
		MandatoryPitDone:         int64(memory.MandatoryPitDone),
		WindSpeed:                float64(memory.WindSpeed),
		WindDirection:            float64(memory.WindDirection),
		IsSetupMenuVisible:       int64(memory.IsSetupMenuVisible),
		MainDisplayIndex:         int64(memory.MainDisplayIndex),
		SecondaryDisplayIndex:    int64(memory.SecondaryDisplayIndex),
		Tc:                       int64(memory.Tc),
		TcCut:                    int64(memory.TcCut),
		EngineMap:                int64(memory.EngineMap),
		Abs:                      int64(memory.Abs),
		FuelXLap:                 float64(memory.FuelXLap),
		RainLights:               int64(memory.RainLights),
		FlashingLights:           int64(memory.FlashingLights),
		LightsStage:              int64(memory.LightsStage),
		ExhaustTemperature:       float64(memory.ExhaustTemperature),
		WiperLV:                  int64(memory.WiperLV),
		DriverStintTotalTimeLeft: int64(memory.DriverStintTotalTimeLeft),
		DriverStintTimeLeft:      int64(memory.DriverStintTimeLeft),
		RainTyres:                int64(memory.RainTyres),
		SessionIndex:             int64(memory.SessionIndex),
		UsedFuel:                 float64(memory.UsedFuel),
		DeltaLapTime:             uint16ToString(memory.DeltaLapTime[:]),
		IDeltaLapTime:            int64(memory.IDeltaLapTime),
		EstimatedLapTime:         uint16ToString(memory.EstimatedLapTime[:]),
		IEstimatedLapTime:        int64(memory.IEstimatedLapTime),
		IsDeltaPositive:          int64(memory.IsDeltaPositive),
		ISplit:                   int64(memory.ISplit),
		IsValidLap:               int64(memory.IsValidLap),
		FuelEstimatedLaps:        float64(memory.FuelEstimatedLaps),
		TrackStatus:              uint16ToString(memory.TrackStatus[:]),
		MissingMandatoryPits:     int64(memory.MissingMandatoryPits),
		Clock:                    float64(memory.Clock),
		DirectionLightsLeft:      int64(memory.DirectionLightsLeft),
		DirectionLightsRight:     int64(memory.DirectionLightsRight),
		GlobalYellow:             int64(memory.GlobalYellow),
		GlobalYellow1:            int64(memory.GlobalYellow1),
		GlobalYellow2:            int64(memory.GlobalYellow2),
		GlobalYellow3:            int64(memory.GlobalYellow3),
		GlobalWhite:              int64(memory.GlobalWhite),
		GlobalGreen:              int64(memory.GlobalGreen),
		GlobalChequered:          int64(memory.GlobalChequered),
		GlobalRed:                int64(memory.GlobalRed),
		MfdTyreSet:               int64(memory.MfdTyreSet),
		MfdFuelToAdd:             float64(memory.MfdFuelToAdd),
		MfdTyrePressureLF:        float64(memory.MfdTyrePressureLF),
		MfdTyrePressureRF:        float64(memory.MfdTyrePressureRF),
		MfdTyrePressureLR:        float64(memory.MfdTyrePressureLR),
		MfdTyrePressureRR:        float64(memory.MfdTyrePressureRR),
		TrackGripStatus:          int64(memory.TrackGripStatus),
		RainIntensity:            int64(memory.RainIntensity),
		RainIntensityIn10min:     int64(memory.RainIntensityIn10min),
		RainIntensityIn30min:     int64(memory.RainIntensityIn30min),
		CurrentTyreSet:           int64(memory.CurrentTyreSet),
		StrategyTyreSet:          int64(memory.StrategyTyreSet),
		GapAhead:                 int64(memory.GapAhead),
		GapBehind:                int64(memory.GapBehind),
	}
}

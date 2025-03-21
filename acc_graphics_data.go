package model

import (
	"fmt"
	"hash/fnv"
	"time"
	"unsafe"
)

type AccGraphicsData struct {
	PacketId                 int32
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
	ReplayTimeMultiplier     float32 // Not used
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

func (memory *AccGraphicsData) Name() string {
	return ACC_GRAPHICS_FILE_NAME
}

func (memory *AccGraphicsData) Path() string {
	return ACC_FILES_PREFIX + memory.Name()
}

func (memory *AccGraphicsData) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d", memory.PacketId)))
	return h.Sum32()
}

func (memory *AccGraphicsData) MapValues(pointer uintptr) {
	newValue := (*AccGraphicsData)(unsafe.Pointer(pointer))
	*memory = *newValue
}

func (memory *AccGraphicsData) CreateMetric() Metrics {
	return &AccGraphicsMetrics{
		Timestamp:                time.Now().UTC(),
		PacketId:                 int64(memory.PacketId),
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
		AarCoordinates:           twoDimensionSliceFloat32To64(floatArray_60_3_ToSlice(memory.AarCoordinates)),
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

package model

import (
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

func (memory *AccGraphicsData) CreateMetric(pointer uintptr) Metrics {
	memory = (*AccGraphicsData)(unsafe.Pointer(pointer))
	return &AccGraphicsMetrics{
		Timestamp:                time.Now().UTC(),
		PacketId:                 memory.PacketId,
		Status:                   memory.Status,
		Session:                  memory.Session,
		CurrentTime:              uint16ToString(memory.CurrentTime[:]),
		LastTime:                 uint16ToString(memory.LastTime[:]),
		BestTime:                 uint16ToString(memory.BestTime[:]),
		Split:                    uint16ToString(memory.Split[:]),
		CompletedLaps:            memory.CompletedLaps,
		Position:                 memory.Position,
		ICurrentTime:             memory.ICurrentTime,
		ILastTime:                memory.ILastTime,
		IBestTime:                memory.IBestTime,
		SessionTimeLeft:          memory.SessionTimeLeft,
		DistanceTraveled:         memory.DistanceTraveled,
		IsInPit:                  memory.IsInPit,
		CurrentSectorIndex:       memory.CurrentSectorIndex,
		LastSectorTime:           memory.LastSectorTime,
		NumberOfLaps:             memory.NumberOfLaps,
		TyreCompound:             uint16ToString(memory.TyreCompound[:]),
		NormalizedCarPosition:    memory.NormalizedCarPosition,
		ActiveCars:               memory.ActiveCars,
		AarCoordinates:           memory.AarCoordinates,
		CarID:                    memory.CarID,
		PlayerCarID:              memory.PlayerCarID,
		PenaltyTime:              memory.PenaltyTime,
		Flag:                     memory.Flag,
		Penalty:                  memory.Penalty,
		IdealLineOn:              memory.IdealLineOn,
		IsInPitLane:              memory.IsInPitLane,
		SurfaceGrip:              memory.SurfaceGrip,
		MandatoryPitDone:         memory.MandatoryPitDone,
		WindSpeed:                memory.WindSpeed,
		WindDirection:            memory.WindDirection,
		IsSetupMenuVisible:       memory.IsSetupMenuVisible,
		MainDisplayIndex:         memory.MainDisplayIndex,
		SecondaryDisplayIndex:    memory.SecondaryDisplayIndex,
		Tc:                       memory.Tc,
		TcCut:                    memory.TcCut,
		EngineMap:                memory.EngineMap,
		Abs:                      memory.Abs,
		FuelXLap:                 memory.FuelXLap,
		RainLights:               memory.RainLights,
		FlashingLights:           memory.FlashingLights,
		LightsStage:              memory.LightsStage,
		ExhaustTemperature:       memory.ExhaustTemperature,
		WiperLV:                  memory.WiperLV,
		DriverStintTotalTimeLeft: memory.DriverStintTotalTimeLeft,
		DriverStintTimeLeft:      memory.DriverStintTimeLeft,
		RainTyres:                memory.RainTyres,
		SessionIndex:             memory.SessionIndex,
		UsedFuel:                 memory.UsedFuel,
		DeltaLapTime:             uint16ToString(memory.DeltaLapTime[:]),
		IDeltaLapTime:            memory.IDeltaLapTime,
		EstimatedLapTime:         uint16ToString(memory.EstimatedLapTime[:]),
		IEstimatedLapTime:        memory.IEstimatedLapTime,
		IsDeltaPositive:          memory.IsDeltaPositive,
		ISplit:                   memory.ISplit,
		IsValidLap:               memory.IsValidLap,
		FuelEstimatedLaps:        memory.FuelEstimatedLaps,
		TrackStatus:              uint16ToString(memory.TrackStatus[:]),
		MissingMandatoryPits:     memory.MissingMandatoryPits,
		Clock:                    memory.Clock,
		DirectionLightsLeft:      memory.DirectionLightsLeft,
		DirectionLightsRight:     memory.DirectionLightsRight,
		GlobalYellow:             memory.GlobalYellow,
		GlobalYellow1:            memory.GlobalYellow1,
		GlobalYellow2:            memory.GlobalYellow2,
		GlobalYellow3:            memory.GlobalYellow3,
		GlobalWhite:              memory.GlobalWhite,
		GlobalGreen:              memory.GlobalGreen,
		GlobalChequered:          memory.GlobalChequered,
		GlobalRed:                memory.GlobalRed,
		MfdTyreSet:               memory.MfdTyreSet,
		MfdFuelToAdd:             memory.MfdFuelToAdd,
		MfdTyrePressureLF:        memory.MfdTyrePressureLF,
		MfdTyrePressureRF:        memory.MfdTyrePressureRF,
		MfdTyrePressureLR:        memory.MfdTyrePressureLR,
		MfdTyrePressureRR:        memory.MfdTyrePressureRR,
		TrackGripStatus:          memory.TrackGripStatus,
		RainIntensity:            memory.RainIntensity,
		RainIntensityIn10min:     memory.RainIntensityIn10min,
		RainIntensityIn30min:     memory.RainIntensityIn30min,
		CurrentTyreSet:           memory.CurrentTyreSet,
		StrategyTyreSet:          memory.StrategyTyreSet,
		GapAhead:                 memory.GapAhead,
		GapBehind:                memory.GapBehind,
	}
}

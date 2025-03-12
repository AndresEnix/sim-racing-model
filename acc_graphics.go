package model

import (
	"fmt"
	"unicode/utf16"
	"unsafe"
)

type AccGraphicsMemory struct {
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
	SecondaryDisplyIndex     int32
	TC                       int32
	TCCUT                    int32
	EngineMap                int32
	ABS                      int32
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

func (memory AccGraphicsMemory) FileName() string {
	return "Local\\acpmf_graphics"
}

func (memory AccGraphicsMemory) Create(pointer uintptr) DataMemoryMapping {
	return (*AccGraphicsMemory)(unsafe.Pointer(pointer))
}

func (memory AccGraphicsMemory) ToMetric() Metric {
	return AccGraphicsMetric{
		PacketId:                 memory.PacketId,
		Status:                   memory.Status,
		Session:                  memory.Session,
		CurrentTime:              string(utf16.Decode(memory.CurrentTime[:])),
		LastTime:                 string(utf16.Decode(memory.LastTime[:])),
		BestTime:                 string(utf16.Decode(memory.BestTime[:])),
		Split:                    string(utf16.Decode(memory.Split[:])),
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
		TyreCompound:             string(utf16.Decode(memory.TyreCompound[:])),
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
		SecondaryDisplyIndex:     memory.SecondaryDisplyIndex,
		TC:                       memory.TC,
		TCCUT:                    memory.TCCUT,
		EngineMap:                memory.EngineMap,
		ABS:                      memory.ABS,
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
		DeltaLapTime:             string(utf16.Decode(memory.DeltaLapTime[:])),
		IDeltaLapTime:            memory.IDeltaLapTime,
		EstimatedLapTime:         string(utf16.Decode(memory.EstimatedLapTime[:])),
		IEstimatedLapTime:        memory.IEstimatedLapTime,
		IsDeltaPositive:          memory.IsDeltaPositive,
		ISplit:                   memory.ISplit,
		IsValidLap:               memory.IsValidLap,
		FuelEstimatedLaps:        memory.FuelEstimatedLaps,
		TrackStatus:              string(utf16.Decode(memory.TrackStatus[:])),
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

type AccGraphicsMetric struct {
	PacketId                 int32          `json:"PacketId"`
	Status                   int32          `json:"Status"`
	Session                  int32          `json:"Session"`
	CurrentTime              string         `json:"CurrentTime"`
	LastTime                 string         `json:"LastTime"`
	BestTime                 string         `json:"BestTime"`
	Split                    string         `json:"Split"`
	CompletedLaps            int32          `json:"CompletedLaps"`
	Position                 int32          `json:"Position"`
	ICurrentTime             int32          `json:"ICurrentTime"`
	ILastTime                int32          `json:"ILastTime"`
	IBestTime                int32          `json:"IBestTime"`
	SessionTimeLeft          float32        `json:"SessionTimeLeft"`
	DistanceTraveled         float32        `json:"DistanceTraveled"`
	IsInPit                  int32          `json:"IsInPit"`
	CurrentSectorIndex       int32          `json:"CurrentSectorIndex"`
	LastSectorTime           int32          `json:"LastSectorTime"`
	NumberOfLaps             int32          `json:"NumberOfLaps"`
	TyreCompound             string         `json:"TyreCompound"`
	NormalizedCarPosition    float32        `json:"NormalizedCarPosition"`
	ActiveCars               int32          `json:"ActiveCars"`
	AarCoordinates           [60][3]float32 `json:"AarCoordinates"`
	CarID                    [60]int32      `json:"CarID"`
	PlayerCarID              int32          `json:"PlayerCarID"`
	PenaltyTime              float32        `json:"PenaltyTime"`
	Flag                     int32          `json:"Flag"`
	Penalty                  int32          `json:"Penalty"`
	IdealLineOn              int32          `json:"IdealLineOn"`
	IsInPitLane              int32          `json:"IsInPitLane"`
	SurfaceGrip              float32        `json:"SurfaceGrip"`
	MandatoryPitDone         int32          `json:"MandatoryPitDone"`
	WindSpeed                float32        `json:"WindSpeed"`
	WindDirection            float32        `json:"WindDirection"`
	IsSetupMenuVisible       int32          `json:"IsSetupMenuVisible"`
	MainDisplayIndex         int32          `json:"MainDisplayIndex"`
	SecondaryDisplyIndex     int32          `json:"SecondaryDisplyIndex"`
	TC                       int32          `json:"TC"`
	TCCUT                    int32          `json:"TCCUT"`
	EngineMap                int32          `json:"EngineMap"`
	ABS                      int32          `json:"ABS"`
	FuelXLap                 float32        `json:"FuelXLap"`
	RainLights               int32          `json:"RainLights"`
	FlashingLights           int32          `json:"FlashingLights"`
	LightsStage              int32          `json:"LightsStage"`
	ExhaustTemperature       float32        `json:"ExhaustTemperature"`
	WiperLV                  int32          `json:"WiperLV"`
	DriverStintTotalTimeLeft int32          `json:"DriverStintTotalTimeLeft"`
	DriverStintTimeLeft      int32          `json:"DriverStintTimeLeft"`
	RainTyres                int32          `json:"RainTyres"`
	SessionIndex             int32          `json:"SessionIndex"`
	UsedFuel                 float32        `json:"UsedFuel"`
	DeltaLapTime             string         `json:"DeltaLapTime"`
	IDeltaLapTime            int32          `json:"IDeltaLapTime"`
	EstimatedLapTime         string         `json:"EstimatedLapTime"`
	IEstimatedLapTime        int32          `json:"IEstimatedLapTime"`
	IsDeltaPositive          int32          `json:"IsDeltaPositive"`
	ISplit                   int32          `json:"ISplit"`
	IsValidLap               int32          `json:"IsValidLap"`
	FuelEstimatedLaps        float32        `json:"FuelEstimatedLaps"`
	TrackStatus              string         `json:"TrackStatus"`
	MissingMandatoryPits     int32          `json:"MissingMandatoryPits"`
	Clock                    float32        `json:"Clock"`
	DirectionLightsLeft      int32          `json:"DirectionLightsLeft"`
	DirectionLightsRight     int32          `json:"DirectionLightsRight"`
	GlobalYellow             int32          `json:"GlobalYellow"`
	GlobalYellow1            int32          `json:"GlobalYellow1"`
	GlobalYellow2            int32          `json:"GlobalYellow2"`
	GlobalYellow3            int32          `json:"GlobalYellow3"`
	GlobalWhite              int32          `json:"GlobalWhite"`
	GlobalGreen              int32          `json:"GlobalGreen"`
	GlobalChequered          int32          `json:"GlobalChequered"`
	GlobalRed                int32          `json:"GlobalRed"`
	MfdTyreSet               int32          `json:"MfdTyreSet"`
	MfdFuelToAdd             float32        `json:"MfdFuelToAdd"`
	MfdTyrePressureLF        float32        `json:"MfdTyrePressureLF"`
	MfdTyrePressureRF        float32        `json:"MfdTyrePressureRF"`
	MfdTyrePressureLR        float32        `json:"MfdTyrePressureLR"`
	MfdTyrePressureRR        float32        `json:"MfdTyrePressureRR"`
	TrackGripStatus          int32          `json:"TrackGripStatus"`
	RainIntensity            int32          `json:"RainIntensity"`
	RainIntensityIn10min     int32          `json:"RainIntensityIn10min"`
	RainIntensityIn30min     int32          `json:"RainIntensityIn30min"`
	CurrentTyreSet           int32          `json:"CurrentTyreSet"`
	StrategyTyreSet          int32          `json:"StrategyTyreSet"`
	GapAhead                 int32          `json:"GapAhead"`
	GapBehind                int32          `json:"GapBehind"`
}

func (metric AccGraphicsMetric) Print() {
	fmt.Println("")
}

package model

import "time"

type AccGraphicsMetrics struct {
	Id                       string         `json:"Id"`
	UserId                   string         `json:"UserId"`
	SessionId                string         `json:"SessionId"`
	Timestamp                time.Time      `json:"Timestamp"`
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
	SecondaryDisplayIndex    int32          `json:"SecondaryDisplayIndex"`
	Tc                       int32          `json:"Tc"`
	TcCut                    int32          `json:"TcCut"`
	EngineMap                int32          `json:"EngineMap"`
	Abs                      int32          `json:"Abs"`
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

func (metric *AccGraphicsMetrics) Game() string {
	return ASSETTO_CORSA_COMPETIZIONE
}

func (metric *AccGraphicsMetrics) Name() string {
	return ACC_GRAPHICS_METRICS_NAME
}

func (metric *AccGraphicsMetrics) DataPoints() []string {
	return getStructFieldNames(*metric)
}

func (metric *AccGraphicsMetrics) AddSessionInfo(userId, sessionId string) {
	metric.UserId = userId
	metric.SessionId = sessionId
}

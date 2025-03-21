package model

import "time"

type AccGraphicsMetrics struct {
	Id                       string      `json:"Id"`
	UserId                   string      `json:"UserId"`
	SessionId                string      `json:"SessionId"`
	Timestamp                time.Time   `json:"Timestamp"`
	PacketId                 int64       `json:"PacketId"`
	Status                   int64       `json:"Status"`
	Session                  int64       `json:"Session"`
	CurrentTime              string      `json:"CurrentTime"`
	LastTime                 string      `json:"LastTime"`
	BestTime                 string      `json:"BestTime"`
	Split                    string      `json:"Split"`
	CompletedLaps            int64       `json:"CompletedLaps"`
	Position                 int64       `json:"Position"`
	ICurrentTime             int64       `json:"ICurrentTime"`
	ILastTime                int64       `json:"ILastTime"`
	IBestTime                int64       `json:"IBestTime"`
	SessionTimeLeft          float64     `json:"SessionTimeLeft"`
	DistanceTraveled         float64     `json:"DistanceTraveled"`
	IsInPit                  int64       `json:"IsInPit"`
	CurrentSectorIndex       int64       `json:"CurrentSectorIndex"`
	LastSectorTime           int64       `json:"LastSectorTime"`
	NumberOfLaps             int64       `json:"NumberOfLaps"`
	TyreCompound             string      `json:"TyreCompound"`
	NormalizedCarPosition    float64     `json:"NormalizedCarPosition"`
	ActiveCars               int64       `json:"ActiveCars"`
	AarCoordinates           [][]float64 `json:"AarCoordinates"`
	CarID                    []int64     `json:"CarID"`
	PlayerCarID              int64       `json:"PlayerCarID"`
	PenaltyTime              float64     `json:"PenaltyTime"`
	Flag                     int64       `json:"Flag"`
	Penalty                  int64       `json:"Penalty"`
	IdealLineOn              int64       `json:"IdealLineOn"`
	IsInPitLane              int64       `json:"IsInPitLane"`
	SurfaceGrip              float64     `json:"SurfaceGrip"`
	MandatoryPitDone         int64       `json:"MandatoryPitDone"`
	WindSpeed                float64     `json:"WindSpeed"`
	WindDirection            float64     `json:"WindDirection"`
	IsSetupMenuVisible       int64       `json:"IsSetupMenuVisible"`
	MainDisplayIndex         int64       `json:"MainDisplayIndex"`
	SecondaryDisplayIndex    int64       `json:"SecondaryDisplayIndex"`
	Tc                       int64       `json:"Tc"`
	TcCut                    int64       `json:"TcCut"`
	EngineMap                int64       `json:"EngineMap"`
	Abs                      int64       `json:"Abs"`
	FuelXLap                 float64     `json:"FuelXLap"`
	RainLights               int64       `json:"RainLights"`
	FlashingLights           int64       `json:"FlashingLights"`
	LightsStage              int64       `json:"LightsStage"`
	ExhaustTemperature       float64     `json:"ExhaustTemperature"`
	WiperLV                  int64       `json:"WiperLV"`
	DriverStintTotalTimeLeft int64       `json:"DriverStintTotalTimeLeft"`
	DriverStintTimeLeft      int64       `json:"DriverStintTimeLeft"`
	RainTyres                int64       `json:"RainTyres"`
	SessionIndex             int64       `json:"SessionIndex"`
	UsedFuel                 float64     `json:"UsedFuel"`
	DeltaLapTime             string      `json:"DeltaLapTime"`
	IDeltaLapTime            int64       `json:"IDeltaLapTime"`
	EstimatedLapTime         string      `json:"EstimatedLapTime"`
	IEstimatedLapTime        int64       `json:"IEstimatedLapTime"`
	IsDeltaPositive          int64       `json:"IsDeltaPositive"`
	ISplit                   int64       `json:"ISplit"`
	IsValidLap               int64       `json:"IsValidLap"`
	FuelEstimatedLaps        float64     `json:"FuelEstimatedLaps"`
	TrackStatus              string      `json:"TrackStatus"`
	MissingMandatoryPits     int64       `json:"MissingMandatoryPits"`
	Clock                    float64     `json:"Clock"`
	DirectionLightsLeft      int64       `json:"DirectionLightsLeft"`
	DirectionLightsRight     int64       `json:"DirectionLightsRight"`
	GlobalYellow             int64       `json:"GlobalYellow"`
	GlobalYellow1            int64       `json:"GlobalYellow1"`
	GlobalYellow2            int64       `json:"GlobalYellow2"`
	GlobalYellow3            int64       `json:"GlobalYellow3"`
	GlobalWhite              int64       `json:"GlobalWhite"`
	GlobalGreen              int64       `json:"GlobalGreen"`
	GlobalChequered          int64       `json:"GlobalChequered"`
	GlobalRed                int64       `json:"GlobalRed"`
	MfdTyreSet               int64       `json:"MfdTyreSet"`
	MfdFuelToAdd             float64     `json:"MfdFuelToAdd"`
	MfdTyrePressureLF        float64     `json:"MfdTyrePressureLF"`
	MfdTyrePressureRF        float64     `json:"MfdTyrePressureRF"`
	MfdTyrePressureLR        float64     `json:"MfdTyrePressureLR"`
	MfdTyrePressureRR        float64     `json:"MfdTyrePressureRR"`
	TrackGripStatus          int64       `json:"TrackGripStatus"`
	RainIntensity            int64       `json:"RainIntensity"`
	RainIntensityIn10min     int64       `json:"RainIntensityIn10min"`
	RainIntensityIn30min     int64       `json:"RainIntensityIn30min"`
	CurrentTyreSet           int64       `json:"CurrentTyreSet"`
	StrategyTyreSet          int64       `json:"StrategyTyreSet"`
	GapAhead                 int64       `json:"GapAhead"`
	GapBehind                int64       `json:"GapBehind"`
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

func (metric *AccGraphicsMetrics) New() Metrics {
	return &AccGraphicsMetrics{}
}

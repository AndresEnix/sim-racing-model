package model

import (
	"reflect"
	"time"
)

// AccGraphicsMetrics is the name of the AccGraphics metrics.
type AccGraphicsMetrics struct {
	ID                       string      `json:"id"`
	UserID                   string      `json:"userId"`
	SessionID                string      `json:"sessionId"`
	Timestamp                time.Time   `json:"timestamp"`
	PacketID                 int64       `json:"packetId"`
	Status                   int64       `json:"status"`
	Session                  int64       `json:"session"`
	CurrentTime              string      `json:"currentTime"`
	LastTime                 string      `json:"lastTime"`
	BestTime                 string      `json:"bestTime"`
	Split                    string      `json:"split"`
	CompletedLaps            int64       `json:"completedLaps"`
	Position                 int64       `json:"position"`
	ICurrentTime             int64       `json:"iCurrentTime"`
	ILastTime                int64       `json:"iLastTime"`
	IBestTime                int64       `json:"iBestTime"`
	SessionTimeLeft          float64     `json:"sessionTimeLeft"`
	DistanceTraveled         float64     `json:"distanceTraveled"`
	IsInPit                  int64       `json:"isInPit"`
	CurrentSectorIndex       int64       `json:"currentSectorIndex"`
	LastSectorTime           int64       `json:"lastSectorTime"`
	NumberOfLaps             int64       `json:"numberOfLaps"`
	TyreCompound             string      `json:"tyreCompound"`
	NormalizedCarPosition    float64     `json:"normalizedCarPosition"`
	ActiveCars               int64       `json:"activeCars"`
	AarCoordinates           [][]float64 `json:"aarCoordinates"`
	CarID                    []int64     `json:"carId"`
	PlayerCarID              int64       `json:"playerCarId"`
	PenaltyTime              float64     `json:"penaltyTime"`
	Flag                     int64       `json:"flag"`
	Penalty                  int64       `json:"penalty"`
	IdealLineOn              int64       `json:"idealLineOn"`
	IsInPitLane              int64       `json:"isInPitLane"`
	SurfaceGrip              float64     `json:"surfaceGrip"`
	MandatoryPitDone         int64       `json:"mandatoryPitDone"`
	WindSpeed                float64     `json:"windSpeed"`
	WindDirection            float64     `json:"windDirection"`
	IsSetupMenuVisible       int64       `json:"isSetupMenuVisible"`
	MainDisplayIndex         int64       `json:"mainDisplayIndex"`
	SecondaryDisplayIndex    int64       `json:"secondaryDisplayIndex"`
	Tc                       int64       `json:"tc"`
	TcCut                    int64       `json:"tcCut"`
	EngineMap                int64       `json:"engineMap"`
	Abs                      int64       `json:"abs"`
	FuelXLap                 float64     `json:"fuelXLap"`
	RainLights               int64       `json:"rainLights"`
	FlashingLights           int64       `json:"flashingLights"`
	LightsStage              int64       `json:"lightsStage"`
	ExhaustTemperature       float64     `json:"exhaustTemperature"`
	WiperLV                  int64       `json:"wiperLv"`
	DriverStintTotalTimeLeft int64       `json:"driverStintTotalTimeLeft"`
	DriverStintTimeLeft      int64       `json:"driverStintTimeLeft"`
	RainTyres                int64       `json:"rainTyres"`
	SessionIndex             int64       `json:"sessionIndex"`
	UsedFuel                 float64     `json:"usedFuel"`
	DeltaLapTime             string      `json:"deltaLapTime"`
	IDeltaLapTime            int64       `json:"iDeltaLapTime"`
	EstimatedLapTime         string      `json:"estimatedLapTime"`
	IEstimatedLapTime        int64       `json:"iEstimatedLapTime"`
	IsDeltaPositive          int64       `json:"isDeltaPositive"`
	ISplit                   int64       `json:"iSplit"`
	IsValidLap               int64       `json:"isValidLap"`
	FuelEstimatedLaps        float64     `json:"fuelEstimatedLaps"`
	TrackStatus              string      `json:"trackStatus"`
	MissingMandatoryPits     int64       `json:"missingMandatoryPits"`
	Clock                    float64     `json:"clock"`
	DirectionLightsLeft      int64       `json:"directionLightsLeft"`
	DirectionLightsRight     int64       `json:"directionLightsRight"`
	GlobalYellow             int64       `json:"globalYellow"`
	GlobalYellow1            int64       `json:"globalYellow1"`
	GlobalYellow2            int64       `json:"globalYellow2"`
	GlobalYellow3            int64       `json:"globalYellow3"`
	GlobalWhite              int64       `json:"globalWhite"`
	GlobalGreen              int64       `json:"globalGreen"`
	GlobalChequered          int64       `json:"globalChequered"`
	GlobalRed                int64       `json:"globalRed"`
	MfdTyreSet               int64       `json:"mfdTyreSet"`
	MfdFuelToAdd             float64     `json:"mfdFuelToAdd"`
	MfdTyrePressureLF        float64     `json:"mfdTyrePressureLf"`
	MfdTyrePressureRF        float64     `json:"mfdTyrePressureRf"`
	MfdTyrePressureLR        float64     `json:"mfdTyrePressureLr"`
	MfdTyrePressureRR        float64     `json:"mfdTyrePressureRr"`
	TrackGripStatus          int64       `json:"trackGripStatus"`
	RainIntensity            int64       `json:"rainIntensity"`
	RainIntensityIn10min     int64       `json:"rainIntensityIn10min"`
	RainIntensityIn30min     int64       `json:"rainIntensityIn30min"`
	CurrentTyreSet           int64       `json:"currentTyreSet"`
	StrategyTyreSet          int64       `json:"strategyTyreSet"`
	GapAhead                 int64       `json:"gapAhead"`
	GapBehind                int64       `json:"gapBehind"`
}

//nolint:funlen
// NewAccGraphicsMetrics creates a new AccGraphicsMetrics struct with default values.
func NewAccGraphicsMetrics() *AccGraphicsMetrics {
	return &AccGraphicsMetrics{
		ID:                       "",
		UserID:                   "",
		SessionID:                "",
		Timestamp:                time.Time{},
		PacketID:                 0,
		Status:                   0,
		Session:                  0,
		CurrentTime:              "",
		LastTime:                 "",
		BestTime:                 "",
		Split:                    "",
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
		TyreCompound:             "",
		NormalizedCarPosition:    0.0,
		ActiveCars:               0,
		AarCoordinates:           make([][]float64, 0),
		CarID:                    make([]int64, 0),
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
		DeltaLapTime:             "",
		IDeltaLapTime:            0,
		EstimatedLapTime:         "",
		IEstimatedLapTime:        0,
		IsDeltaPositive:          0,
		ISplit:                   0,
		IsValidLap:               0,
		FuelEstimatedLaps:        0.0,
		TrackStatus:              "",
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

// New creates a new AccGraphicsMetrics struct with default values.
func (metric *AccGraphicsMetrics) New() Metrics {
	return NewAccGraphicsMetrics()
}

// Game returns the name of the game for the AccGraphicsMetrics metric.
func (metric *AccGraphicsMetrics) Game() string {
	return AssettoCorsaCompetizione
}

// Name returns the name of the AccGraphicsMetrics metric.
func (metric *AccGraphicsMetrics) Name() string {
	return AccGraphicsMetricsName
}

// DataPoints returns the field names of the AccGraphicsMetrics struct.
func (metric *AccGraphicsMetrics) DataPoints() []string {
	v := reflect.ValueOf(metric)
	t := v.Type()

	fieldNames := make([]string, 0, v.NumField())
	for i := range v.NumField() {
		fieldNames = append(fieldNames, t.Field(i).Name)
	}

	return fieldNames
}

// AddSessionInfo adds the user ID and session ID to the metric.
func (metric *AccGraphicsMetrics) AddSessionInfo(userID, sessionID string) {
	metric.UserID = userID
	metric.SessionID = sessionID
}

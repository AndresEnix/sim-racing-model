package model

import (
	"reflect"
	"time"
)

// AccPhysicsMetrics is the name of the AccPhysics metrics.
type AccPhysicsMetrics struct {
	ID                  string      `json:"id"`
	UserID              string      `json:"userId"`
	SessionID           string      `json:"sessionId"`
	Timestamp           time.Time   `json:"timestamp"`
	PacketID            int64       `json:"packetId"`
	Gas                 float64     `json:"gas"`
	Brake               float64     `json:"brake"`
	Fuel                float64     `json:"fuel"`
	Gear                int64       `json:"gear"`
	RPM                 int64       `json:"rpm"`
	SteerAngle          float64     `json:"steerAngle"`
	SpeedKmh            float64     `json:"speedKmh"`
	Velocity            []float64   `json:"velocity"`
	AccG                []float64   `json:"accG"`
	WheelSlip           []float64   `json:"wheelSlip"`
	WheelPressure       []float64   `json:"wheelPressure"`
	WheelAngularSpeed   []float64   `json:"wheelAngularSpeed"`
	TyreCoreTemperature []float64   `json:"tyreCoreTemperature"`
	SuspensionTravel    []float64   `json:"suspensionTravel"`
	TC                  float64     `json:"tc"`
	Heading             float64     `json:"heading"`
	Pitch               float64     `json:"pitch"`
	Roll                float64     `json:"roll"`
	CarDamage           []float64   `json:"carDamage"`
	PitLimiterOn        int64       `json:"pitLimiterOn"`
	Abs                 float64     `json:"abs"`
	AutoShifterOn       int64       `json:"autoShifterOn"`
	TurboBoost          float64     `json:"turboBoost"`
	AirTemp             float64     `json:"airTemp"`
	RoadTemp            float64     `json:"roadTemp"`
	LocalAngularVel     []float64   `json:"localAngularVel"`
	FinalFF             float64     `json:"finalFf"`
	BrakeTemp           []float64   `json:"brakeTemp"`
	Clutch              float64     `json:"clutch"`
	IsAIControlled      int64       `json:"isAiControlled"`
	TyreContactPoint    [][]float64 `json:"tyreContactPoint"`
	TyreContactNormal   [][]float64 `json:"tyreContactNormal"`
	TyreContactHeading  [][]float64 `json:"tyreContactHeading"`
	BrakeBias           float64     `json:"brakeBias"`
	LocalVelocity       []float64   `json:"localVelocity"`
	SlipRatio           []float64   `json:"slipRatio"`
	SlipAngle           []float64   `json:"slipAngle"`
	WaterTemp           float64     `json:"waterTemp"`
	BrakePressure       []float64   `json:"brakePressure"`
	FrontBrakeCompound  int64       `json:"frontBrakeCompound"`
	RearBrakeCompound   int64       `json:"rearBrakeCompound"`
	PadLife             []float64   `json:"padLife"`
	DiscLife            []float64   `json:"discLife"`
	IgnitionOn          int64       `json:"ignitionOn"`
	StarterEngineOn     int64       `json:"starterEngineOn"`
	IsEngineRunning     int64       `json:"isEngineRunning"`
	KerbVibration       float64     `json:"kerbVibration"`
	SlipVibrations      float64     `json:"slipVibrations"`
	GVibrations         float64     `json:"gVibrations"`
	AbsVibrations       float64     `json:"absVibrations"`
}

// NewAccPhysicsMetrics creates a new AccPhysicsMetrics struct with default values.
func NewAccPhysicsMetrics() *AccPhysicsMetrics {
	return &AccPhysicsMetrics{
		ID:                  "",
		UserID:              "",
		SessionID:           "",
		Timestamp:           time.Time{},
		PacketID:            0,
		Gas:                 0.0,
		Brake:               0.0,
		Fuel:                0.0,
		Gear:                0,
		RPM:                 0,
		SteerAngle:          0.0,
		SpeedKmh:            0.0,
		Velocity:            make([]float64, 0),
		AccG:                make([]float64, 0),
		WheelSlip:           make([]float64, 0),
		WheelPressure:       make([]float64, 0),
		WheelAngularSpeed:   make([]float64, 0),
		TyreCoreTemperature: make([]float64, 0),
		SuspensionTravel:    make([]float64, 0),
		TC:                  0.0,
		Heading:             0.0,
		Pitch:               0.0,
		Roll:                0.0,
		CarDamage:           make([]float64, 0),
		PitLimiterOn:        0,
		Abs:                 0.0,
		AutoShifterOn:       0,
		TurboBoost:          0.0,
		AirTemp:             0.0,
		RoadTemp:            0.0,
		LocalAngularVel:     make([]float64, 0),
		FinalFF:             0.0,
		BrakeTemp:           make([]float64, 0),
		Clutch:              0.0,
		IsAIControlled:      0,
		TyreContactPoint:    make([][]float64, 0),
		TyreContactNormal:   make([][]float64, 0),
		TyreContactHeading:  make([][]float64, 0),
		BrakeBias:           0.0,
		LocalVelocity:       make([]float64, 0),
		SlipRatio:           make([]float64, 0),
		SlipAngle:           make([]float64, 0),
		WaterTemp:           0.0,
		BrakePressure:       make([]float64, 0),
		FrontBrakeCompound:  0,
		RearBrakeCompound:   0,
		PadLife:             make([]float64, 0),
		DiscLife:            make([]float64, 0),
		IgnitionOn:          0,
		StarterEngineOn:     0,
		IsEngineRunning:     0,
		KerbVibration:       0.0,
		SlipVibrations:      0.0,
		GVibrations:         0.0,
		AbsVibrations:       0.0,
	}
}

// New creates a new AccPhysicsMetrics struct with default values.
func (metric *AccPhysicsMetrics) New() Metrics {
	return NewAccPhysicsMetrics()
}

// Game returns the name of the game for the AccPhysicsMetrics metric.
func (metric *AccPhysicsMetrics) Game() string {
	return AssettoCorsaCompetizione
}

// Name returns the name of the AccPhysicsMetrics metric.
func (metric *AccPhysicsMetrics) Name() string {
	return AccPhysicsMetricsName
}

// DataPoints returns the field names of the AccPhysicsMetrics struct.
func (metric *AccPhysicsMetrics) DataPoints() []string {
	v := reflect.ValueOf(metric)
	t := v.Type()

	fieldNames := make([]string, 0, v.NumField())
	for i := range v.NumField() {
		fieldNames = append(fieldNames, t.Field(i).Name)
	}

	return fieldNames
}

// AddSessionInfo adds the user ID and session ID to the metric.
func (metric *AccPhysicsMetrics) AddSessionInfo(userID, sessionID string) {
	metric.UserID = userID
	metric.SessionID = sessionID
}

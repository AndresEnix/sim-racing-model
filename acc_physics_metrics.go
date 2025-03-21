package model

import "time"

type AccPhysicsMetrics struct {
	Id                  string      `json:"Id"`
	UserId              string      `json:"UserId"`
	SessionId           string      `json:"SessionId"`
	Timestamp           time.Time   `json:"Timestamp"`
	PacketId            int64       `json:"PacketId"`
	Gas                 float64     `json:"Gas"`
	Brake               float64     `json:"Brake"`
	Fuel                float64     `json:"Fuel"`
	Gear                int64       `json:"Gear"`
	RPM                 int64       `json:"RPM"`
	SteerAngle          float64     `json:"SteerAngle"`
	SpeedKmh            float64     `json:"SpeedKmh"`
	Velocity            []float64   `json:"Velocity"`
	AccG                []float64   `json:"AccG"`
	WheelSlip           []float64   `json:"WheelSlip"`
	WheelPressure       []float64   `json:"WheelPressure"`
	WheelAngularSpeed   []float64   `json:"WheelAngularSpeed"`
	TyreCoreTemperature []float64   `json:"TyreCoreTemperature"`
	SuspensionTravel    []float64   `json:"SuspensionTravel"`
	TC                  float64     `json:"TC"`
	Heading             float64     `json:"Heading"`
	Pitch               float64     `json:"Pitch"`
	Roll                float64     `json:"Roll"`
	CarDamage           []float64   `json:"CarDamage"`
	PitLimiterOn        int64       `json:"PitLimiterOn"`
	Abs                 float64     `json:"Abs"`
	AutoShifterOn       int64       `json:"AutoShifterOn"`
	TurboBoost          float64     `json:"TurboBoost"`
	AirTemp             float64     `json:"AirTemp"`
	RoadTemp            float64     `json:"RoadTemp"`
	LocalAngularVel     []float64   `json:"LocalAngularVel"`
	FinalFF             float64     `json:"FinalFF"`
	BrakeTemp           []float64   `json:"BrakeTemp"`
	Clutch              float64     `json:"Clutch"`
	IsAIControlled      int64       `json:"IsAIControlled"`
	TyreContactPoint    [][]float64 `json:"TyreContactPoint"`
	TyreContactNormal   [][]float64 `json:"TyreContactNormal"`
	TyreContactHeading  [][]float64 `json:"TyreContactHeading"`
	BrakeBias           float64     `json:"BrakeBias"`
	LocalVelocity       []float64   `json:"LocalVelocity"`
	SlipRatio           []float64   `json:"SlipRatio"`
	SlipAngle           []float64   `json:"SlipAngle"`
	WaterTemp           float64     `json:"WaterTemp"`
	BrakePressure       []float64   `json:"BrakePressure"`
	FrontBrakeCompound  int64       `json:"FrontBrakeCompound"`
	RearBrakeCompound   int64       `json:"RearBrakeCompound"`
	PadLife             []float64   `json:"PadLife"`
	DiscLife            []float64   `json:"DiscLife"`
	IgnitionOn          int64       `json:"IgnitionOn"`
	StarterEngineOn     int64       `json:"StarterEngineOn"`
	IsEngineRunning     int64       `json:"IsEngineRunning"`
	KerbVibration       float64     `json:"KerbVibration"`
	SlipVibrations      float64     `json:"SlipVibrations"`
	GVibrations         float64     `json:"GVibrations"`
	AbsVibrations       float64     `json:"AbsVibrations"`
}

func (metric *AccPhysicsMetrics) Game() string {
	return ASSETTO_CORSA_COMPETIZIONE
}

func (metric *AccPhysicsMetrics) Name() string {
	return ACC_PHYSICS_METRICS_NAME
}

func (metric *AccPhysicsMetrics) DataPoints() []string {
	return getStructFieldNames(*metric)
}

func (metric *AccPhysicsMetrics) AddSessionInfo(userId, sessionId string) {
	metric.UserId = userId
	metric.SessionId = sessionId
}

func (metric *AccPhysicsMetrics) New() Metrics {
	return &AccPhysicsMetrics{}
}

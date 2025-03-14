package model

import "time"

type AccPhysicsMetrics struct {
	Id                  string        `json:"Id"`
	UserId              string        `json:"UserId"`
	SessionId           string        `json:"SessionId"`
	Timestamp           time.Time     `json:"Timestamp"`
	PacketId            int32         `json:"PacketId"`
	Gas                 float32       `json:"Gas"`
	Brake               float32       `json:"Brake"`
	Fuel                float32       `json:"Fuel"`
	Gear                int32         `json:"Gear"`
	RPM                 int32         `json:"RPM"`
	SteerAngle          float32       `json:"SteerAngle"`
	SpeedKmh            float32       `json:"SpeedKmh"`
	Velocity            [3]float32    `json:"Velocity"`
	AccG                [3]float32    `json:"AccG"`
	WheelSlip           [4]float32    `json:"WheelSlip"`
	WheelPressure       [4]float32    `json:"WheelPressure"`
	WheelAngularSpeed   [4]float32    `json:"WheelAngularSpeed"`
	TyreCoreTemperature [4]float32    `json:"TyreCoreTemperature"`
	SuspensionTravel    [4]float32    `json:"SuspensionTravel"`
	TC                  float32       `json:"TC"`
	Heading             float32       `json:"Heading"`
	Pitch               float32       `json:"Pitch"`
	Roll                float32       `json:"Roll"`
	CarDamage           [5]float32    `json:"CarDamage"`
	PitLimiterOn        int32         `json:"PitLimiterOn"`
	Abs                 float32       `json:"Abs"`
	AutoShifterOn       int32         `json:"AutoShifterOn"`
	TurboBoost          float32       `json:"TurboBoost"`
	AirTemp             float32       `json:"AirTemp"`
	RoadTemp            float32       `json:"RoadTemp"`
	LocalAngularVel     [3]float32    `json:"LocalAngularVel"`
	FinalFF             float32       `json:"FinalFF"`
	BrakeTemp           [4]float32    `json:"BrakeTemp"`
	Clutch              float32       `json:"Clutch"`
	IsAIControlled      int32         `json:"IsAIControlled"`
	TyreContactPoint    [4][3]float32 `json:"TyreContactPoint"`
	TyreContactNormal   [4][3]float32 `json:"TyreContactNormal"`
	TyreContactHeading  [4][3]float32 `json:"TyreContactHeading"`
	BrakeBias           float32       `json:"BrakeBias"`
	LocalVelocity       [3]float32    `json:"LocalVelocity"`
	SlipRatio           [4]float32    `json:"SlipRatio"`
	SlipAngle           [4]float32    `json:"SlipAngle"`
	WaterTemp           float32       `json:"WaterTemp"`
	BrakePressure       [4]float32    `json:"BrakePressure"`
	FrontBrakeCompound  int32         `json:"FrontBrakeCompound"`
	RearBrakeCompound   int32         `json:"RearBrakeCompound"`
	PadLife             [4]float32    `json:"PadLife"`
	DiscLife            [4]float32    `json:"DiscLife"`
	IgnitionOn          int32         `json:"IgnitionOn"`
	StarterEngineOn     int32         `json:"StarterEngineOn"`
	IsEngineRunning     int32         `json:"IsEngineRunning"`
	KerbVibration       float32       `json:"KerbVibration"`
	SlipVibrations      float32       `json:"SlipVibrations"`
	GVibrations         float32       `json:"GVibrations"`
	AbsVibrations       float32       `json:"AbsVibrations"`
}

func (metric *AccPhysicsMetrics) Game() string {
	return ASSETTO_CORSA_COMPETIZIONE
}

func (metric *AccPhysicsMetrics) Name() string {
	return ACC_PHYSICS_METRICS_NAME
}

func (metric *AccPhysicsMetrics) DataPoints() []string {
	return getStructFieldNames(metric)
}

func (metric *AccPhysicsMetrics) AddSessionInfo(userId, sessionId string) {
	metric.UserId = userId
	metric.SessionId = sessionId
}

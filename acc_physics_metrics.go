package model

import (
	"reflect"
	"time"
)

// AccPhysicsMetrics is the name of the AccPhysics metrics.
type AccPhysicsMetrics struct {
	ID                  uint        `gorm:"primaryKey;column:id"                                    json:"id"`
	UserID              string      `gorm:"column:user_id"                                          json:"userId"`
	SessionID           string      `gorm:"column:session_id"                                       json:"sessionId"`
	Timestamp           time.Time   `gorm:"column:timestamp"                                        json:"timestamp"`
	PacketID            int64       `gorm:"column:packet_id"                                        json:"packetId"`
	Gas                 float64     `gorm:"column:gas"                                              json:"gas"`
	Brake               float64     `gorm:"column:brake"                                            json:"brake"`
	Fuel                float64     `gorm:"column:fuel"                                             json:"fuel"`
	Gear                int64       `gorm:"column:gear"                                             json:"gear"`
	RPM                 int64       `gorm:"column:rpm"                                              json:"rpm"`
	SteerAngle          float64     `gorm:"column:steer_angle"                                      json:"steerAngle"`
	SpeedKmh            float64     `gorm:"column:speed_kmh"                                        json:"speedKmh"`
	Velocity            []float64   `gorm:"type:jsonb;serializer:json;column:velocity"              json:"velocity"`
	AccG                []float64   `gorm:"type:jsonb;serializer:json;column:acc_g"                 json:"accG"`
	WheelSlip           []float64   `gorm:"type:jsonb;serializer:json;column:wheel_slip"            json:"wheelSlip"`
	WheelPressure       []float64   `gorm:"type:jsonb;serializer:json;column:wheel_pressure"        json:"wheelPressure"`
	WheelAngularSpeed   []float64   `gorm:"type:jsonb;serializer:json;column:wheel_angular_speed"   json:"wheelAngularSpeed"`
	TyreCoreTemperature []float64   `gorm:"type:jsonb;serializer:json;column:tyre_core_temperature" json:"tyreCoreTemperature"`
	SuspensionTravel    []float64   `gorm:"type:jsonb;serializer:json;column:suspension_travel"     json:"suspensionTravel"`
	TC                  float64     `gorm:"column:tc"                                               json:"tc"`
	Heading             float64     `gorm:"column:heading"                                          json:"heading"`
	Pitch               float64     `gorm:"column:pitch"                                            json:"pitch"`
	Roll                float64     `gorm:"column:roll"                                             json:"roll"`
	CarDamage           []float64   `gorm:"type:jsonb;serializer:json;column:car_damage"            json:"carDamage"`
	PitLimiterOn        int64       `gorm:"column:pit_limiter_on"                                   json:"pitLimiterOn"`
	Abs                 float64     `gorm:"column:abs"                                              json:"abs"`
	AutoShifterOn       int64       `gorm:"column:auto_shifter_on"                                  json:"autoShifterOn"`
	TurboBoost          float64     `gorm:"column:turbo_boost"                                      json:"turboBoost"`
	AirTemp             float64     `gorm:"column:air_temp"                                         json:"airTemp"`
	RoadTemp            float64     `gorm:"column:road_temp"                                        json:"roadTemp"`
	LocalAngularVel     []float64   `gorm:"type:jsonb;serializer:json;column:local_angular_vel"     json:"localAngularVel"`
	FinalFF             float64     `gorm:"column:final_ff"                                         json:"finalFf"`
	BrakeTemp           []float64   `gorm:"type:jsonb;serializer:json;column:brake_temp"            json:"brakeTemp"`
	Clutch              float64     `gorm:"column:clutch"                                           json:"clutch"`
	IsAIControlled      int64       `gorm:"column:is_ai_controlled"                                 json:"isAiControlled"`
	TyreContactPoint    [][]float64 `gorm:"type:jsonb;serializer:json;column:tyre_contact_point"    json:"tyreContactPoint"`
	TyreContactNormal   [][]float64 `gorm:"type:jsonb;serializer:json;column:tyre_contact_normal"   json:"tyreContactNormal"`
	TyreContactHeading  [][]float64 `gorm:"type:jsonb;serializer:json;column:tyre_contact_heading"  json:"tyreContactHeading"`
	BrakeBias           float64     `gorm:"column:brake_bias"                                       json:"brakeBias"`
	LocalVelocity       []float64   `gorm:"type:jsonb;serializer:json;column:local_velocity"        json:"localVelocity"`
	SlipRatio           []float64   `gorm:"type:jsonb;serializer:json;column:slip_ratio"            json:"slipRatio"`
	SlipAngle           []float64   `gorm:"type:jsonb;serializer:json;column:slip_angle"            json:"slipAngle"`
	WaterTemp           float64     `gorm:"column:water_temp"                                       json:"waterTemp"`
	BrakePressure       []float64   `gorm:"type:jsonb;serializer:json;column:brake_pressure"        json:"brakePressure"`
	FrontBrakeCompound  int64       `gorm:"column:front_brake_compound"                             json:"frontBrakeCompound"`
	RearBrakeCompound   int64       `gorm:"column:rear_brake_compound"                              json:"rearBrakeCompound"`
	PadLife             []float64   `gorm:"type:jsonb;serializer:json;column:pad_life"              json:"padLife"`
	DiscLife            []float64   `gorm:"type:jsonb;serializer:json;column:disc_life"             json:"discLife"`
	IgnitionOn          int64       `gorm:"column:ignition_on"                                      json:"ignitionOn"`
	StarterEngineOn     int64       `gorm:"column:starter_engine_on"                                json:"starterEngineOn"`
	IsEngineRunning     int64       `gorm:"column:is_engine_running"                                json:"isEngineRunning"`
	KerbVibration       float64     `gorm:"column:kerb_vibration"                                   json:"kerbVibration"`
	SlipVibrations      float64     `gorm:"column:slip_vibrations"                                  json:"slipVibrations"`
	GVibrations         float64     `gorm:"column:g_vibrations"                                     json:"gVibrations"`
	AbsVibrations       float64     `gorm:"column:abs_vibrations"                                   json:"absVibrations"`
}

// NewAccPhysicsMetrics creates a new AccPhysicsMetrics struct with default values.
func NewAccPhysicsMetrics() *AccPhysicsMetrics {
	return &AccPhysicsMetrics{
		ID:                  0,
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
	metricType := reflect.TypeOf(metric)

	if metricType.Kind() == reflect.Ptr {
		metricType = metricType.Elem()
	}

	if metricType.Kind() != reflect.Struct {
		return []string{}
	}

	var names []string

	for i := range metricType.NumField() {
		field := metricType.Field(i)

		if field.IsExported() {
			names = append(names, field.Name)
		}
	}

	return names
}

// TableName returns the name of the DB table for the AccPhysicsMetrics metric.
func (metric *AccPhysicsMetrics) TableName() string {
	return "acc_physics_metrics"
}

// AddSessionInfo adds the user ID and session ID to the metric.
func (metric *AccPhysicsMetrics) AddSessionInfo(userID, sessionID string) {
	metric.UserID = userID
	metric.SessionID = sessionID
}

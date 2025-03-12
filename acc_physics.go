package model

import (
	"fmt"
	"unsafe"
)

type AccPhysicsMemory struct {
	PacketId            int32
	Gas                 float32
	Brake               float32
	Fuel                float32
	Gear                int32
	RPM                 int32
	SteerAngle          float32
	SpeedKmh            float32
	Velocity            [3]float32
	AccG                [3]float32
	WheelSlip           [4]float32
	WheelLoad           [4]float32 // Not used
	WheelPressure       [4]float32
	WheelAngularSpeed   [4]float32
	TyreWear            [4]float32 // Not used
	TyreDirtyLevel      [4]float32 // Not used
	TyreCoreTemperature [4]float32
	CamberRAD           [4]float32 // Not used
	SuspensionTravel    [4]float32
	DRS                 float32 // Not used
	TC                  float32
	Heading             float32
	Pitch               float32
	Roll                float32
	CgHeight            float32 // Not used
	CarDamage           [5]float32
	NumberOfTyresOut    int32 // Not used
	PitLimiterOn        int32
	ABS                 float32
	KersCharge          float32 // Not used
	KersInput           float32 // Not used
	AutoshifterOn       int32
	RideHeight          [2]float32 // Not used
	TurboBoost          float32
	Ballast             float32 // Not used
	AirDensity          float32 // Not used
	AirTemp             float32
	RoadTemp            float32
	LocalAngularVel     [3]float32
	FinalFF             float32
	PerformanceMeter    float32 // Not used
	EngineBrake         int32   // Not used
	ErsRecoveryLevel    int32   // Not used
	ErsPowerLevel       int32   // Not used
	ErsHeatCharging     int32   // Not used
	ErsIsCharging       int32   // Not used
	KersCurrentKJ       float32 // Not used
	DrsAvailable        int32   // Not used
	DrsEnabled          int32   // Not used
	BrakeTemp           [4]float32
	Clutch              float32
	TyreTempI           [4]float32 // Not used
	TyreTempM           [4]float32 // Not used
	TyreTempO           [4]float32 // Not used
	IsAIControlled      int32
	TyreContactPoint    [4][3]float32
	TyreContactNormal   [4][3]float32
	TyreContactHeading  [4][3]float32
	BrakeBias           float32
	LocalVelocity       [3]float32
	P2PActivation       int32      // Not used
	P2PStatus           int32      // Not used
	CurrentMaxRpm       float32    // Not used
	MZ                  [4]float32 // Not used
	FX                  [4]float32 // Not used
	FY                  [4]float32 // Not used
	SlipRatio           [4]float32
	SlipAngle           [4]float32
	TcinAction          int32      // Not used
	AbsInAction         int32      // Not used
	SuspensionDamage    [4]float32 // Not used
	TyreTemp            [4]float32 // Not used
	WaterTemp           float32
	BrakePressure       [4]float32
	FrontBrakeCompound  int32
	RearBrakeCompound   int32
	PadLife             [4]float32
	DiscLife            [4]float32
	IgnitionOn          int32
	StarterEngineOn     int32
	IsEngineRunning     int32
	KerbVibration       float32
	SlipVibrations      float32
	GVibrations         float32
	AbsVibrations       float32
}

func (memory AccPhysicsMemory) FileName() string {
	return "Local\\acpmf_physics"
}

func (memory AccPhysicsMemory) Create(pointer uintptr) DataMemoryMapping {
	return (*AccPhysicsMemory)(unsafe.Pointer(pointer))
}

func (memory AccPhysicsMemory) ToMetric() Metric {
	return PhysicsMetric{
		PacketId:            memory.PacketId,
		Gas:                 memory.Gas,
		Brake:               memory.Brake,
		Fuel:                memory.Fuel,
		Gear:                memory.Gear,
		RPM:                 memory.RPM,
		SteerAngle:          memory.SteerAngle,
		SpeedKmh:            memory.SpeedKmh,
		Velocity:            memory.Velocity,
		AccG:                memory.AccG,
		WheelSlip:           memory.WheelSlip,
		WheelPressure:       memory.WheelPressure,
		WheelAngularSpeed:   memory.WheelAngularSpeed,
		TyreCoreTemperature: memory.TyreCoreTemperature,
		SuspensionTravel:    memory.SuspensionTravel,
		TC:                  memory.TC,
		Heading:             memory.Heading,
		Pitch:               memory.Pitch,
		Roll:                memory.Roll,
		CarDamage:           memory.CarDamage,
		PitLimiterOn:        memory.PitLimiterOn,
		ABS:                 memory.ABS,
		AutoshifterOn:       memory.AutoshifterOn,
		TurboBoost:          memory.TurboBoost,
		AirTemp:             memory.AirTemp,
		RoadTemp:            memory.RoadTemp,
		LocalAngularVel:     memory.LocalAngularVel,
		FinalFF:             memory.FinalFF,
		BrakeTemp:           memory.BrakeTemp,
		Clutch:              memory.Clutch,
		IsAIControlled:      memory.IsAIControlled,
		TyreContactPoint:    memory.TyreContactPoint,
		TyreContactNormal:   memory.TyreContactNormal,
		TyreContactHeading:  memory.TyreContactHeading,
		BrakeBias:           memory.BrakeBias,
		LocalVelocity:       memory.LocalVelocity,
		SlipRatio:           memory.SlipRatio,
		SlipAngle:           memory.SlipAngle,
		WaterTemp:           memory.WaterTemp,
		BrakePressure:       memory.BrakePressure,
		FrontBrakeCompound:  memory.FrontBrakeCompound,
		RearBrakeCompound:   memory.RearBrakeCompound,
		PadLife:             memory.PadLife,
		DiscLife:            memory.DiscLife,
		IgnitionOn:          memory.IgnitionOn,
		StarterEngineOn:     memory.StarterEngineOn,
		IsEngineRunning:     memory.IsEngineRunning,
		KerbVibration:       memory.KerbVibration,
		SlipVibrations:      memory.SlipVibrations,
		GVibrations:         memory.GVibrations,
		AbsVibrations:       memory.AbsVibrations,
	}
}

type PhysicsMetric struct {
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
	ABS                 float32       `json:"ABS"`
	AutoshifterOn       int32         `json:"AutoshifterOn"`
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

func (metric PhysicsMetric) Print() {
	fmt.Println("")
}

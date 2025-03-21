package model

import (
	"fmt"
	"hash/fnv"
	"time"
	"unsafe"
)

type AccPhysicsData struct {
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
	Abs                 float32
	KersCharge          float32 // Not used
	KersInput           float32 // Not used
	AutoShifterOn       int32
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
	TcInAction          int32      // Not used
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

func (memory *AccPhysicsData) Name() string {
	return ACC_PHYSICS_FILE_NAME
}

func (memory *AccPhysicsData) Path() string {
	return ACC_FILES_PREFIX + memory.Name()
}

func (memory *AccPhysicsData) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d", memory.PacketId)))
	return h.Sum32()
}

func (memory *AccPhysicsData) MapValues(pointer uintptr) {
	newValue := (*AccPhysicsData)(unsafe.Pointer(pointer))
	*memory = *newValue
}

func (memory *AccPhysicsData) CreateMetric() Metrics {
	return &AccPhysicsMetrics{
		Timestamp:           time.Now().UTC(),
		PacketId:            int64(memory.PacketId),
		Gas:                 float64(memory.Gas),
		Brake:               float64(memory.Brake),
		Fuel:                float64(memory.Fuel),
		Gear:                int64(memory.Gear),
		RPM:                 int64(memory.RPM),
		SteerAngle:          float64(memory.SteerAngle),
		SpeedKmh:            float64(memory.SpeedKmh),
		Velocity:            oneDimensionSliceFloat32To64(memory.Velocity[:]),
		AccG:                oneDimensionSliceFloat32To64(memory.AccG[:]),
		WheelSlip:           oneDimensionSliceFloat32To64(memory.WheelSlip[:]),
		WheelPressure:       oneDimensionSliceFloat32To64(memory.WheelPressure[:]),
		WheelAngularSpeed:   oneDimensionSliceFloat32To64(memory.WheelAngularSpeed[:]),
		TyreCoreTemperature: oneDimensionSliceFloat32To64(memory.TyreCoreTemperature[:]),
		SuspensionTravel:    oneDimensionSliceFloat32To64(memory.SuspensionTravel[:]),
		TC:                  float64(memory.TC),
		Heading:             float64(memory.Heading),
		Pitch:               float64(memory.Pitch),
		Roll:                float64(memory.Roll),
		CarDamage:           oneDimensionSliceFloat32To64(memory.CarDamage[:]),
		PitLimiterOn:        int64(memory.PitLimiterOn),
		Abs:                 float64(memory.Abs),
		AutoShifterOn:       int64(memory.AutoShifterOn),
		TurboBoost:          float64(memory.TurboBoost),
		AirTemp:             float64(memory.AirTemp),
		RoadTemp:            float64(memory.RoadTemp),
		LocalAngularVel:     oneDimensionSliceFloat32To64(memory.LocalAngularVel[:]),
		FinalFF:             float64(memory.FinalFF),
		BrakeTemp:           oneDimensionSliceFloat32To64(memory.BrakeTemp[:]),
		Clutch:              float64(memory.Clutch),
		IsAIControlled:      int64(memory.IsAIControlled),
		TyreContactPoint:    twoDimensionSliceFloat32To64(floatArray_4_3_ToSlice(memory.TyreContactPoint)),
		TyreContactNormal:   twoDimensionSliceFloat32To64(floatArray_4_3_ToSlice(memory.TyreContactNormal)),
		TyreContactHeading:  twoDimensionSliceFloat32To64(floatArray_4_3_ToSlice(memory.TyreContactHeading)),
		BrakeBias:           float64(memory.BrakeBias),
		LocalVelocity:       oneDimensionSliceFloat32To64(memory.LocalVelocity[:]),
		SlipRatio:           oneDimensionSliceFloat32To64(memory.SlipRatio[:]),
		SlipAngle:           oneDimensionSliceFloat32To64(memory.SlipAngle[:]),
		WaterTemp:           float64(memory.WaterTemp),
		BrakePressure:       oneDimensionSliceFloat32To64(memory.BrakePressure[:]),
		FrontBrakeCompound:  int64(memory.FrontBrakeCompound),
		RearBrakeCompound:   int64(memory.RearBrakeCompound),
		PadLife:             oneDimensionSliceFloat32To64(memory.PadLife[:]),
		DiscLife:            oneDimensionSliceFloat32To64(memory.DiscLife[:]),
		IgnitionOn:          int64(memory.IgnitionOn),
		StarterEngineOn:     int64(memory.StarterEngineOn),
		IsEngineRunning:     int64(memory.IsEngineRunning),
		KerbVibration:       float64(memory.KerbVibration),
		SlipVibrations:      float64(memory.SlipVibrations),
		GVibrations:         float64(memory.GVibrations),
		AbsVibrations:       float64(memory.AbsVibrations),
	}
}

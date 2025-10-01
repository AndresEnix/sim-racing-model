package model

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"math"
	"time"
	"unsafe"
)

// AccPhysicsData represents the structure of the ACC graphics data shared memory.
type AccPhysicsData struct {
	PacketID            int32
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
	WheelLoad           [4]float32 // Not used.
	WheelPressure       [4]float32
	WheelAngularSpeed   [4]float32
	TyreWear            [4]float32 // Not used.
	TyreDirtyLevel      [4]float32 // Not used.
	TyreCoreTemperature [4]float32
	CamberRAD           [4]float32 // Not used.
	SuspensionTravel    [4]float32
	DRS                 float32 // Not used.
	TC                  float32
	Heading             float32
	Pitch               float32
	Roll                float32
	CgHeight            float32 // Not used.
	CarDamage           [5]float32
	NumberOfTyresOut    int32 // Not used.
	PitLimiterOn        int32
	Abs                 float32
	KersCharge          float32 // Not used.
	KersInput           float32 // Not used.
	AutoShifterOn       int32
	RideHeight          [2]float32 // Not used.
	TurboBoost          float32
	Ballast             float32 // Not used.
	AirDensity          float32 // Not used.
	AirTemp             float32
	RoadTemp            float32
	LocalAngularVel     [3]float32
	FinalFF             float32
	PerformanceMeter    float32 // Not used.
	EngineBrake         int32   // Not used.
	ErsRecoveryLevel    int32   // Not used.
	ErsPowerLevel       int32   // Not used.
	ErsHeatCharging     int32   // Not used.
	ErsIsCharging       int32   // Not used.
	KersCurrentKJ       float32 // Not used.
	DrsAvailable        int32   // Not used.
	DrsEnabled          int32   // Not used.
	BrakeTemp           [4]float32
	Clutch              float32
	TyreTempI           [4]float32 // Not used.
	TyreTempM           [4]float32 // Not used.
	TyreTempO           [4]float32 // Not used.
	IsAIControlled      int32
	TyreContactPoint    [4][3]float32
	TyreContactNormal   [4][3]float32
	TyreContactHeading  [4][3]float32
	BrakeBias           float32
	LocalVelocity       [3]float32
	P2PActivation       int32      // Not used.
	P2PStatus           int32      // Not used.
	CurrentMaxRpm       float32    // Not used.
	MZ                  [4]float32 // Not used.
	FX                  [4]float32 // Not used.
	FY                  [4]float32 // Not used.
	SlipRatio           [4]float32
	SlipAngle           [4]float32
	TcInAction          int32      // Not used.
	AbsInAction         int32      // Not used.
	SuspensionDamage    [4]float32 // Not used.
	TyreTemp            [4]float32 // Not used.
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

//nolint:funlen
// NewAccPhysicsData creates a new instance of AccPhysicsData with default values.
func NewAccPhysicsData() *AccPhysicsData {
	return &AccPhysicsData{
		PacketID:            0,
		Gas:                 0.0,
		Brake:               0.0,
		Fuel:                0.0,
		Gear:                0,
		RPM:                 0,
		SteerAngle:          0.0,
		SpeedKmh:            0.0,
		Velocity:            [3]float32{},
		AccG:                [3]float32{},
		WheelSlip:           [4]float32{},
		WheelLoad:           [4]float32{},
		WheelPressure:       [4]float32{},
		WheelAngularSpeed:   [4]float32{},
		TyreWear:            [4]float32{},
		TyreDirtyLevel:      [4]float32{},
		TyreCoreTemperature: [4]float32{},
		CamberRAD:           [4]float32{},
		SuspensionTravel:    [4]float32{},
		DRS:                 0.0,
		TC:                  0.0,
		Heading:             0.0,
		Pitch:               0.0,
		Roll:                0.0,
		CgHeight:            0.0,
		CarDamage:           [5]float32{},
		NumberOfTyresOut:    0,
		PitLimiterOn:        0,
		Abs:                 0.0,
		KersCharge:          0.0,
		KersInput:           0.0,
		AutoShifterOn:       0,
		RideHeight:          [2]float32{},
		TurboBoost:          0.0,
		Ballast:             0.0,
		AirDensity:          0.0,
		AirTemp:             0.0,
		RoadTemp:            0.0,
		LocalAngularVel:     [3]float32{},
		FinalFF:             0.0,
		PerformanceMeter:    0.0,
		EngineBrake:         0,
		ErsRecoveryLevel:    0,
		ErsPowerLevel:       0,
		ErsHeatCharging:     0,
		ErsIsCharging:       0,
		KersCurrentKJ:       0.0,
		DrsAvailable:        0,
		DrsEnabled:          0,
		BrakeTemp:           [4]float32{},
		Clutch:              0.0,
		TyreTempI:           [4]float32{},
		TyreTempM:           [4]float32{},
		TyreTempO:           [4]float32{},
		IsAIControlled:      0,
		TyreContactPoint:    [4][3]float32{},
		TyreContactNormal:   [4][3]float32{},
		TyreContactHeading:  [4][3]float32{},
		BrakeBias:           0.0,
		LocalVelocity:       [3]float32{},
		P2PActivation:       0,
		P2PStatus:           0,
		CurrentMaxRpm:       0.0,
		MZ:                  [4]float32{},
		FX:                  [4]float32{},
		FY:                  [4]float32{},
		SlipRatio:           [4]float32{},
		SlipAngle:           [4]float32{},
		TcInAction:          0,
		AbsInAction:         0,
		SuspensionDamage:    [4]float32{},
		TyreTemp:            [4]float32{},
		WaterTemp:           0.0,
		BrakePressure:       [4]float32{},
		FrontBrakeCompound:  0,
		RearBrakeCompound:   0,
		PadLife:             [4]float32{},
		DiscLife:            [4]float32{},
		IgnitionOn:          0,
		StarterEngineOn:     0,
		IsEngineRunning:     0,
		KerbVibration:       0.0,
		SlipVibrations:      0.0,
		GVibrations:         0.0,
		AbsVibrations:       0.0,
	}
}

// Name returns the name of the shared memory file for ACC physics data.
func (memory *AccPhysicsData) Name() string {
	return AccPhysicsFileName
}

// Path returns the full path to the shared memory file for ACC physics data.
func (memory *AccPhysicsData) Path() string {
	return AccFilesPrefix + memory.Name()
}

// ReadFrequency returns the frequency at which the ACC physics data should be read.
func (memory *AccPhysicsData) ReadFrequency() time.Duration {
	frequencyMs := getUint64Env(memory.Name()+ReadFrequencySuffix, AccPhysicsDefaultReadFrequency)
	if frequencyMs > uint64(MaxAllowedMs) {
		log.Printf(
			"WARNING: Duration value (%d ms) from %s overflows time.Duration. Capping at MaxDuration.\n",
			frequencyMs,
			memory.Name()+ReadFrequencySuffix,
		)

		return time.Duration(math.MaxInt64)
	}

	return time.Duration(frequencyMs) * time.Millisecond
}

// Hash generates a hash based on the PacketID of the ACC physics data.
func (memory *AccPhysicsData) Hash() uint32 {
	hasher := fnv.New32a()

	_, err := fmt.Fprintf(hasher, "%d", memory.PacketID)
	if err != nil {
		log.Printf("WARNING: Failed to generate hash for %s: %s\n", memory.Name(), err.Error())

		return 0
	}

	return hasher.Sum32()
}

//nolint:gosec
// MapValues maps the values from a memory pointer to the AccPhysicsData struct.
func (memory *AccPhysicsData) MapValues(pointer uintptr) {
	newValue := (*AccPhysicsData)(unsafe.Pointer(pointer)) //nolint:govet
	*memory = *newValue
}

//nolint:funlen
// CreateMetricsJSON creates JSON bytes from the current AccPhysicsData.
func (memory *AccPhysicsData) CreateMetricsJSON() []byte {
	accPhysicsMetrics := &AccPhysicsMetrics{
		ID:                  "",
		UserID:              "",
		SessionID:           "",
		Timestamp:           time.Now().UTC(),
		PacketID:            int64(memory.PacketID),
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
		TyreContactPoint:    twoDimensionSliceFloat32To64(floatArray4_3ToSlice(memory.TyreContactPoint)),
		TyreContactNormal:   twoDimensionSliceFloat32To64(floatArray4_3ToSlice(memory.TyreContactNormal)),
		TyreContactHeading:  twoDimensionSliceFloat32To64(floatArray4_3ToSlice(memory.TyreContactHeading)),
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

	metrics, err := json.Marshal(accPhysicsMetrics)
	if err != nil {
		log.Println("failed to marshal AccGraphicsData to JSON: %w", err)

		return nil
	}

	return metrics
}

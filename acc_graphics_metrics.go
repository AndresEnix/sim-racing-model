package model

import (
	"reflect"
	"time"
)

// AccGraphicsMetrics is the name of the AccGraphics metrics.
type AccGraphicsMetrics struct {
	ID                       uint        `gorm:"primaryKey;column:id"                              json:"id"`
	UserID                   string      `gorm:"column:user_id"                                    json:"userId"`
	SessionID                string      `gorm:"column:session_id"                                 json:"sessionId"`
	Timestamp                time.Time   `gorm:"column:timestamp"                                  json:"timestamp"`
	PacketID                 int64       `gorm:"column:packet_id"                                  json:"packetId"`
	Status                   int64       `gorm:"column:status"                                     json:"status"`
	Session                  int64       `gorm:"column:session"                                    json:"session"`
	CurrentTime              string      `gorm:"column:current_time"                               json:"currentTime"`
	LastTime                 string      `gorm:"column:last_time"                                  json:"lastTime"`
	BestTime                 string      `gorm:"column:best_time"                                  json:"bestTime"`
	Split                    string      `gorm:"column:split"                                      json:"split"`
	CompletedLaps            int64       `gorm:"column:completed_laps"                             json:"completedLaps"`
	Position                 int64       `gorm:"column:position"                                   json:"position"`
	ICurrentTime             int64       `gorm:"column:i_current_time"                             json:"iCurrentTime"`
	ILastTime                int64       `gorm:"column:i_last_time"                                json:"iLastTime"`
	IBestTime                int64       `gorm:"column:i_best_time"                                json:"iBestTime"`
	SessionTimeLeft          float64     `gorm:"column:session_time_left"                          json:"sessionTimeLeft"`
	DistanceTraveled         float64     `gorm:"column:distance_traveled"                          json:"distanceTraveled"`
	IsInPit                  int64       `gorm:"column:is_in_pit"                                  json:"isInPit"`
	CurrentSectorIndex       int64       `gorm:"column:current_sector_index"                       json:"currentSectorIndex"`
	LastSectorTime           int64       `gorm:"column:last_sector_time"                           json:"lastSectorTime"`
	NumberOfLaps             int64       `gorm:"column:number_of_laps"                             json:"numberOfLaps"`
	TyreCompound             string      `gorm:"column:tyre_compound"                              json:"tyreCompound"`
	NormalizedCarPosition    float64     `gorm:"column:normalized_car_position"                    json:"normalizedCarPosition"`
	ActiveCars               int64       `gorm:"column:active_cars"                                json:"activeCars"`
	AarCoordinates           [][]float64 `gorm:"type:jsonb;serializer:json;column:aar_coordinates" json:"aarCoordinates"`
	CarID                    []int64     `gorm:"type:jsonb;serializer:json;column:car_id"          json:"carId"`
	PlayerCarID              int64       `gorm:"column:player_car_id"                              json:"playerCarId"`
	PenaltyTime              float64     `gorm:"column:penalty_time"                               json:"penaltyTime"`
	Flag                     int64       `gorm:"column:flag"                                       json:"flag"`
	Penalty                  int64       `gorm:"column:penalty"                                    json:"penalty"`
	IdealLineOn              int64       `gorm:"column:ideal_line_on"                              json:"idealLineOn"`
	IsInPitLane              int64       `gorm:"column:is_in_pit_lane"                             json:"isInPitLane"`
	SurfaceGrip              float64     `gorm:"column:surface_grip"                               json:"surfaceGrip"`
	MandatoryPitDone         int64       `gorm:"column:mandatory_pit_done"                         json:"mandatoryPitDone"`
	WindSpeed                float64     `gorm:"column:wind_speed"                                 json:"windSpeed"`
	WindDirection            float64     `gorm:"column:wind_direction"                             json:"windDirection"`
	IsSetupMenuVisible       int64       `gorm:"column:is_setup_menu_visible"                      json:"isSetupMenuVisible"`
	MainDisplayIndex         int64       `gorm:"column:main_display_index"                         json:"mainDisplayIndex"`
	SecondaryDisplayIndex    int64       `gorm:"column:secondary_display_index"                    json:"secondaryDisplayIndex"`
	Tc                       int64       `gorm:"column:tc"                                         json:"tc"`
	TcCut                    int64       `gorm:"column:tc_cut"                                     json:"tcCut"`
	EngineMap                int64       `gorm:"column:engine_map"                                 json:"engineMap"`
	Abs                      int64       `gorm:"column:abs"                                        json:"abs"`
	FuelXLap                 float64     `gorm:"column:fuel_x_lap"                                 json:"fuelXLap"`
	RainLights               int64       `gorm:"column:rain_lights"                                json:"rainLights"`
	FlashingLights           int64       `gorm:"column:flashing_lights"                            json:"flashingLights"`
	LightsStage              int64       `gorm:"column:lights_stage"                               json:"lightsStage"`
	ExhaustTemperature       float64     `gorm:"column:exhaust_temperature"                        json:"exhaustTemperature"`
	WiperLV                  int64       `gorm:"column:wiper_lv"                                   json:"wiperLv"`
	DriverStintTotalTimeLeft int64       `gorm:"column:driver_stint_total_time_left"               json:"driverStintTotalTimeLeft"`
	DriverStintTimeLeft      int64       `gorm:"column:driver_stint_time_left"                     json:"driverStintTimeLeft"`
	RainTyres                int64       `gorm:"column:rain_tyres"                                 json:"rainTyres"`
	SessionIndex             int64       `gorm:"column:session_index"                              json:"sessionIndex"`
	UsedFuel                 float64     `gorm:"column:used_fuel"                                  json:"usedFuel"`
	DeltaLapTime             string      `gorm:"column:delta_lap_time"                             json:"deltaLapTime"`
	IDeltaLapTime            int64       `gorm:"column:i_delta_lap_time"                           json:"iDeltaLapTime"`
	EstimatedLapTime         string      `gorm:"column:estimated_lap_time"                         json:"estimatedLapTime"`
	IEstimatedLapTime        int64       `gorm:"column:i_estimated_lap_time"                       json:"iEstimatedLapTime"`
	IsDeltaPositive          int64       `gorm:"column:is_delta_positive"                          json:"isDeltaPositive"`
	ISplit                   int64       `gorm:"column:i_split"                                    json:"iSplit"`
	IsValidLap               int64       `gorm:"column:is_valid_lap"                               json:"isValidLap"`
	FuelEstimatedLaps        float64     `gorm:"column:fuel_estimated_laps"                        json:"fuelEstimatedLaps"`
	TrackStatus              string      `gorm:"column:track_status"                               json:"trackStatus"`
	MissingMandatoryPits     int64       `gorm:"column:missing_mandatory_pits"                     json:"missingMandatoryPits"`
	Clock                    float64     `gorm:"column:clock"                                      json:"clock"`
	DirectionLightsLeft      int64       `gorm:"column:direction_lights_left"                      json:"directionLightsLeft"`
	DirectionLightsRight     int64       `gorm:"column:direction_lights_right"                     json:"directionLightsRight"`
	GlobalYellow             int64       `gorm:"column:global_yellow"                              json:"globalYellow"`
	GlobalYellow1            int64       `gorm:"column:global_yellow1"                             json:"globalYellow1"`
	GlobalYellow2            int64       `gorm:"column:global_yellow2"                             json:"globalYellow2"`
	GlobalYellow3            int64       `gorm:"column:global_yellow3"                             json:"globalYellow3"`
	GlobalWhite              int64       `gorm:"column:global_white"                               json:"globalWhite"`
	GlobalGreen              int64       `gorm:"column:global_green"                               json:"globalGreen"`
	GlobalChequered          int64       `gorm:"column:global_chequered"                           json:"globalChequered"`
	GlobalRed                int64       `gorm:"column:global_red"                                 json:"globalRed"`
	MfdTyreSet               int64       `gorm:"column:mfd_tyre_set"                               json:"mfdTyreSet"`
	MfdFuelToAdd             float64     `gorm:"column:mfd_fuel_to_add"                            json:"mfdFuelToAdd"`
	MfdTyrePressureLF        float64     `gorm:"column:mfd_tyre_pressure_lf"                       json:"mfdTyrePressureLf"`
	MfdTyrePressureRF        float64     `gorm:"column:mfd_tyre_pressure_rf"                       json:"mfdTyrePressureRf"`
	MfdTyrePressureLR        float64     `gorm:"column:mfd_tyre_pressure_lr"                       json:"mfdTyrePressureLr"`
	MfdTyrePressureRR        float64     `gorm:"column:mfd_tyre_pressure_rr"                       json:"mfdTyrePressureRr"`
	TrackGripStatus          int64       `gorm:"column:track_grip_status"                          json:"trackGripStatus"`
	RainIntensity            int64       `gorm:"column:rain_intensity"                             json:"rainIntensity"`
	RainIntensityIn10min     int64       `gorm:"column:rain_intensity_in_10min"                    json:"rainIntensityIn10min"`
	RainIntensityIn30min     int64       `gorm:"column:rain_intensity_in_30min"                    json:"rainIntensityIn30min"`
	CurrentTyreSet           int64       `gorm:"column:current_tyre_set"                           json:"currentTyreSet"`
	StrategyTyreSet          int64       `gorm:"column:strategy_tyre_set"                          json:"strategyTyreSet"`
	GapAhead                 int64       `gorm:"column:gap_ahead"                                  json:"gapAhead"`
	GapBehind                int64       `gorm:"column:gap_behind"                                 json:"gapBehind"`
}

// NewAccGraphicsMetrics creates a new AccGraphicsMetrics struct with default values.
func NewAccGraphicsMetrics() *AccGraphicsMetrics {
	return &AccGraphicsMetrics{
		ID:                       0,
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

// TableName returns the name of the DB table for the AccGraphicsMetrics metric.
func (metric *AccGraphicsMetrics) TableName() string {
	return "acc_graphics_metrics"
}

// AddSessionInfo adds the user ID and session ID to the metric.
func (metric *AccGraphicsMetrics) AddSessionInfo(userID, sessionID string) {
	metric.UserID = userID
	metric.SessionID = sessionID
}

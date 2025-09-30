package model

import (
	"math"
	"time"
)

// General constants.
const (
	ReadFrequencySuffix = "_read_frequency"
	MaxAllowedMs        = math.MaxInt64 / int64(time.Millisecond)
)

// Game constants.
const (
	AssettoCorsaCompetizione = "acc"
)

// Shared memory constants.
const (
	AccFilesPrefix                  = "Local\\"
	AccPhysicsFileName              = "acpmf_physics"
	AccGraphicsFileName             = "acpmf_graphics"
	AccStaticFileName               = "acpmf_static"
	AccPhysicsDefaultReadFrequency  = 1000
	AccGraphicsDefaultReadFrequency = 1000
	AccStaticDefaultReadFrequency   = 60000
)

// Metric constants.
const (
	AccPhysicsMetricsName  = "acpmf_physics"
	AccGraphicsMetricsName = "acpmf_graphics"
	AccStaticMetricsName   = "acpmf_static"
)

func GetGameMemoryMapping() map[string][]SharedMemoryData {
	return map[string][]SharedMemoryData{
		AssettoCorsaCompetizione: {
			NewAccPhysicsData(),
			NewAccGraphicsData(),
			NewAccStaticData(),
		},
	}
}

func GetGameMetricsMapping() map[string][]Metrics {
	return map[string][]Metrics{
		AssettoCorsaCompetizione: {
			NewAccPhysicsMetrics(),
			NewAccGraphicsMetrics(),
			NewAccStaticMetrics(),
		},
	}
}

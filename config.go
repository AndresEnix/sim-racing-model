package model


import (
	"math"
	"time"
)


const (
	// ReadFrequencySuffix is the suffix for environment variables that specify read frequency in milliseconds.
	ReadFrequencySuffix = "_read_frequency"
	// MaxAllowedMs is the maximum allowed milliseconds to prevent overflow when converting to time.Duration.
	MaxAllowedMs        = math.MaxInt64 / int64(time.Millisecond)
	// AssettoCorsaCompetizione is the fame identifiers and related constants for Assetto Corsa Competizione.
	AssettoCorsaCompetizione = "acc"
	// AccFilesPrefix is the shared memory file prefix.
	AccFilesPrefix                  = "Local\\"
	// AccPhysicsFileName file name.
	AccPhysicsFileName              = "acpmf_physics"
	// AccGraphicsFileName file name.
	AccGraphicsFileName             = "acpmf_graphics"
	// AccStaticFileName file name.
	AccStaticFileName               = "acpmf_static"
	// AccPhysicsDefaultReadFrequency is the default read frequencies in milliseconds for ACC shared memory files.
	AccPhysicsDefaultReadFrequency  = 1000
	// AccGraphicsDefaultReadFrequency is the default read frequencies in milliseconds for ACC shared memory files.
	AccGraphicsDefaultReadFrequency = 1000
	// AccStaticDefaultReadFrequency is the default read frequencies in milliseconds for ACC shared memory files.
	AccStaticDefaultReadFrequency   = 60000
	// AccPhysicsMetricsName metric name.
	AccPhysicsMetricsName  = "acpmf_physics"
	// AccGraphicsMetricsName metric name.
	AccGraphicsMetricsName = "acpmf_graphics"
	// AccStaticMetricsName metric name.
	AccStaticMetricsName   = "acpmf_static"
)


// GetGameMemoryMapping returns a mapping of game IDs to their corresponding shared memory data configurations.
func GetGameMemoryMapping() map[string][]SharedMemoryData {
	return map[string][]SharedMemoryData{
		AssettoCorsaCompetizione: {
			NewAccPhysicsData(),
			NewAccGraphicsData(),
			NewAccStaticData(),
		},
	}
}


// GetGameMetricsMapping returns a mapping of game IDs to their corresponding metrics configurations.
func GetGameMetricsMapping() map[string][]Metrics {
	return map[string][]Metrics{
		AssettoCorsaCompetizione: {
			NewAccPhysicsMetrics(),
			NewAccGraphicsMetrics(),
			NewAccStaticMetrics(),
		},
	}
}

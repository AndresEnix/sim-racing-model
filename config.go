package model

// General constants
const (
	READ_FREQUENCY_SUFFIX      = "_read_frequency"
)

// Game constants
const (
	ASSETTO_CORSA_COMPETIZIONE = "acc"
)

// Shared memory constants
const (
	ACC_FILES_PREFIX       = "Local\\"
	ACC_PHYSICS_FILE_NAME  = "acpmf_physics"
	ACC_GRAPHICS_FILE_NAME = "acpmf_graphics"
	ACC_STATIC_FILE_NAME   = "acpmf_static"
	ACC_PHYSICS_DEFAULT_READ_FREQUENCY = 1000
	ACC_GRAPHICS_DEFAULT_READ_FREQUENCY = 1000
	ACC_STATIC_DEFAULT_READ_FREQUENCY = 60000
)

// Metric constants
const (
	ACC_PHYSICS_METRICS_NAME  = "acpmf_physics"
	ACC_GRAPHICS_METRICS_NAME = "acpmf_graphics"
	ACC_STATIC_METRICS_NAME   = "acpmf_static"
)

// Collections
var (
	gameMemoryMapping = map[string][]SharedMemoryData{
		ASSETTO_CORSA_COMPETIZIONE: {
			&AccPhysicsData{},
			&AccGraphicsData{},
			&AccStaticData{},
		},
	}
	gameMetricsMapping = map[string][]Metrics{
		ASSETTO_CORSA_COMPETIZIONE: {
			&AccPhysicsMetrics{},
			&AccGraphicsMetrics{},
			&AccStaticMetrics{},
		},
	}
)

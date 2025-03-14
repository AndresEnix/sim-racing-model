package model

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
			AccPhysicsData{},
			AccGraphicsData{},
			AccStaticData{},
		},
	}
	gameMetricsMapping = map[string][]Metrics{
		ASSETTO_CORSA_COMPETIZIONE: {
			AccPhysicsMetrics{},
			AccGraphicsMetrics{},
			AccStaticMetrics{},
		},
	}
)

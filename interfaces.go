package model

type SharedMemoryData interface {
	Name() string
	Path() string
	ToMetrics() Metrics
	Create(pointer uintptr) SharedMemoryData
}

type Metrics interface {
	Game() string
	Name() string
	DataPoints() []string
}

package model

import "time"

// SharedMemoryData defines the interface for shared memory data configurations.
type SharedMemoryData interface {
	Name() string
	Path() string
	ReadFrequency() time.Duration
	Hash() uint32
	MapValues(pointer uintptr)
	CreateMetricsJSON() Metrics
}

// Metrics defines the interface for metrics configurations.
type Metrics interface {
	Game() string
	Name() string
	DataPoints() []string
	AddSessionInfo(userID, sessionID string)
	New() Metrics
}

package model


import "time"


// SharedMemoryData defines the interface for shared memory data configurations.
type SharedMemoryData interface {
	Name() string
	Path() string
	ReadFrequency() time.Duration
	MapValues(pointer uintptr)
	Hash() uint32
}


// Metrics defines the interface for metrics configurations.
type Metrics interface {
	Game() string
	Name() string
	DataPoints() []string
	AddSessionInfo(userID, sessionID string)
}

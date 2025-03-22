package model

import "time"

type SharedMemoryData interface {
	Name() string
	Path() string
	ReadFrequency() time.Duration
	MapValues(pointer uintptr)
	CreateMetric() Metrics
	Hash() uint32
}

type Metrics interface {
	Game() string
	Name() string
	DataPoints() []string
	AddSessionInfo(userId, sessionId string)
	New() Metrics
}

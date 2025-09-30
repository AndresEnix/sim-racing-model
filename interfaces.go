package model

import "time"

type SharedMemoryData interface {
	Name() string
	Path() string
	ReadFrequency() time.Duration
	MapValues(pointer uintptr)
	Hash() uint32
}

type Metrics interface {
	Game() string
	Name() string
	DataPoints() []string
	AddSessionInfo(userID, sessionID string)
}

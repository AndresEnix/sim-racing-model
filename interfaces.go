package model

type SharedMemoryData interface {
	Name() string
	Path() string
	CreateMetric(pointer uintptr) Metrics
}

type Metrics interface {
	Game() string
	Name() string
	DataPoints() []string
	AddSessionInfo(userId, sessionId string)
}

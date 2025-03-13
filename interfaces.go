package model

type Metric interface {
	GetFields() []string
}

type DataMemoryMapping interface {
	GetFileName() string
	GetFilePath() string
	Create(pointer uintptr) DataMemoryMapping
	ToMetric() Metric
}

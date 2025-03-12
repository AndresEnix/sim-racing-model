package model

type Metric interface {
	Print()
}

type DataMemoryMapping interface {
	GetFileName() string
	GetFilePath() string
	Create(pointer uintptr) DataMemoryMapping
	ToMetric() Metric
}

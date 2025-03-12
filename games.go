package games

type Metric interface {
    Print()
}

type DataMemoryMapping interface {
	FileName() string
	Create(pointer uintptr) DataMemoryMapping
	ToMetric() Metric
}

package games

type DataMemoryMapping interface {
	FileName() string
	Create(pointer uintptr) DataMemoryMapping
	ToMetric() Metric
}


type Metric interface {
    Print()
}


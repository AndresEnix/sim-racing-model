package model

import (
	"strings"
)

const (
	ASSETTO_CORSA_COMPETIZIONE = "ACC"
)

var (
	gameMemoryMapping = map[string][]DataMemoryMapping{
		ASSETTO_CORSA_COMPETIZIONE: {
			AccPhysicsMemory{},
			AccGraphicsMemory{},
			AccStaticMemory{},
		},
	}
	memoryFimeMetricMapping = map[string]Metric{
		AccPhysicsMemory{}.GetFileName():  AccGraphicsMetric{},
		AccGraphicsMemory{}.GetFileName(): AccGraphicsMetric{},
		AccStaticMemory{}.GetFileName():   AccStaticMetric{},
	}
)

func GetMemoryFiles(gameId string) []DataMemoryMapping {
	return gameMemoryMapping[strings.ToUpper(gameId)]
}

func GetMemoryGames() []string {
	games := make([]string, 0, len(gameMemoryMapping))
	for game := range gameMemoryMapping {
		games = append(games, game)
	}
	return games
}

func GetMetricByMemoryFileName(fileName string) Metric {
	return memoryFimeMetricMapping[fileName]
}

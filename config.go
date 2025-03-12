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
			AccGraphicsMemory{},
			AccPhysicsMemory{},
			AccStaticMemory{},
		},
	}
	AvailableGames = []string{
		ASSETTO_CORSA_COMPETIZIONE,
	}
)

func GetMemoryFiles(gameId string) []DataMemoryMapping {
	return gameMemoryMapping[strings.ToUpper(gameId)]
}

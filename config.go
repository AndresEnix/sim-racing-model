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

package model

import (
	"reflect"
	"strings"
	"unicode/utf16"
)

// Method exposed to the module users
func GameMemoryFiles(gameId string) []SharedMemoryData {
	return gameMemoryMapping[strings.ToLower(gameId)]
}

func GameMetrics(gameId string) []Metrics {
	return gameMetricsMapping[strings.ToLower(gameId)]
}

func GamesUsingSharedMemory() []string {
	games := make([]string, 0, len(gameMemoryMapping))
	for game := range gameMemoryMapping {
		games = append(games, game)
	}
	return games
}

// Internal module methods
func uint16ToString(data []uint16) string {
	cleanData := trimTrailingNulls(data)
	return string(utf16.Decode(cleanData))
}

func trimTrailingNulls(data []uint16) []uint16 {
	for i, v := range data {
		if v == 0 {
			return data[:i]
		}
	}
	return data
}

func getStructFieldNames(instance any) []string {
	v := reflect.ValueOf(instance)
	t := v.Type()

	var fieldNames []string
	for i := range v.NumField() {
		fieldNames = append(fieldNames, t.Field(i).Name)
	}
	return fieldNames
}

package model

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf16"
)

// Exposed methods
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
func getUint64Env(key string, defaultValue uint64) uint64 {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		return defaultValue
	}

	return value
}

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

func floatArray_60_3_ToSlice(arr [60][3]float32) [][]float32 {
	result := make([][]float32, len(arr))
	for i := range arr {
		result[i] = arr[i][:]
	}
	return result
}

func floatArray_4_3_ToSlice(arr [4][3]float32) [][]float32 {
	result := make([][]float32, len(arr))
	for i := range arr {
		result[i] = arr[i][:]
	}
	return result
}

func twoDimensionSliceFloat32To64(input [][]float32) [][]float64 {
	output := make([][]float64, len(input)) // Create outer slice

	for i, row := range input {
		output[i] = oneDimensionSliceFloat32To64(row)
	}

	return output
}

func oneDimensionSliceFloat32To64(input []float32) []float64 {
	output := make([]float64, len(input)) // Create a new slice of float64
	for i, val := range input {
		output[i] = float64(val)
	}
	return output
}

func oneDimensionSliceInt32To64(input []int32) []int64 {
	output := make([]int64, len(input)) // Create a new slice of float64
	for i, val := range input {
		output[i] = int64(val)
	}
	return output
}

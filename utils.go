package model

import (
	"os"
	"strconv"
	"strings"
	"unicode/utf16"
)

// GameMemoryFiles returns the shared memory file mappings for a specific game ID.
func GameMemoryFiles(gameID string) []SharedMemoryData {
	return GetGameMemoryMapping()[strings.ToLower(gameID)]
}

// GameMetrics returns the metrics mappings for a specific game ID.
func GameMetrics(gameID string) []Metrics {
	return GetGameMetricsMapping()[strings.ToLower(gameID)]
}

// GamesUsingSharedMemory returns a list of all games using shared memory.
func GamesUsingSharedMemory() []string {
	games := make([]string, 0, len(GetGameMemoryMapping()))
	for game := range GetGameMemoryMapping() {
		games = append(games, game)
	}

	return games
}

// Helper function to get an environment variable as a uint64 with a default value.
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

// Converts a slice of uint16 (representing UTF-16 encoded characters) to a string,
// trimming any trailing null characters.
func uint16ToString(data []uint16) string {
	cleanData := trimTrailingNulls(data)

	return string(utf16.Decode(cleanData))
}

// Trims trailing null characters (0 values) from a slice of uint16.
func trimTrailingNulls(data []uint16) []uint16 {
	for i, v := range data {
		if v == 0 {
			return data[:i]
		}
	}

	return data
}

// Converts a 2D array of float32 to a slice of slices of float32.
func floatArray60_3ToSlice(arr [60][3]float32) [][]float32 {
	result := make([][]float32, len(arr))
	for i := range arr {
		result[i] = arr[i][:]
	}

	return result
}

// Converts a 2D array of float32 to a slice of slices of float32.
func floatArray4_3ToSlice(arr [4][3]float32) [][]float32 {
	result := make([][]float32, len(arr))
	for i := range arr {
		result[i] = arr[i][:]
	}

	return result
}

// Converts a 2D slice of float32 to a 2D slice of float64.
func twoDimensionSliceFloat32To64(input [][]float32) [][]float64 {
	output := make([][]float64, len(input)) // Create outer slice.

	for i, row := range input {
		output[i] = oneDimensionSliceFloat32To64(row)
	}

	return output
}

// Converts a 1D slice of float32 to a 1D slice of float64.
func oneDimensionSliceFloat32To64(input []float32) []float64 {
	output := make([]float64, len(input)) // Create a new slice of float64.
	for i, val := range input {
		output[i] = float64(val)
	}

	return output
}

// Converts a 2D slice of int32 to a 2D slice of int64.
func oneDimensionSliceInt32To64(input []int32) []int64 {
	output := make([]int64, len(input)) // Create a new slice of float64.
	for i, val := range input {
		output[i] = int64(val)
	}

	return output
}

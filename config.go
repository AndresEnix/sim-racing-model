package model

var (
	gameMapping = map[string][]DataMemoryMapping{
		"ACC": {
			AccGraphicsMemory{},
			AccPhysicsMemory{},
			AccStaticMemory{},
		},
	}
)

func GetMemoryFiles(gameId string) []DataMemoryMapping {
	return gameMapping[gameId]
}

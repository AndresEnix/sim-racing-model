package model

func GetMemoryFiles(fileName string, targetStruct any) map[string][]DataMemoryMapping {
	return map[string][]DataMemoryMapping{
		"ACC": {
			AccGraphicsMemory{},
			AccPhysicsMemory{},
			AccStaticMemory{},
		},
	}
}

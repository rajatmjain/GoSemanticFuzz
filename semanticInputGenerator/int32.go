package semanticInputGenerator

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
)

func GenerateInt32(minimum int32,maximum int32)(int32){
	var generatedInt32 int32
	int32Schema := goSemanticFuzz.Int32Schema{Minimum: minimum,Maximum: minimum}
	int32Fuzzer := goSemanticFuzz.New().Funcs(int32Schema.CustomInt32FuzzFunc())
	int32Fuzzer.Fuzz(&generatedInt32)
	return generatedInt32
}
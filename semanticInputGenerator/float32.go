package semanticInputGenerator

import (
	goSemanticFuzz "GoSemanticFuzz/gosemanticfuzz"
)

func GenerateFloat32(minimum float32, maximum float32, precision uint)(float32){
	var generatedFloat32 float32
	float32Schema := goSemanticFuzz.Float32Schema{Minimum: minimum, Maximum: maximum,Precision: precision}
	float32Fuzzer := goSemanticFuzz.New().Funcs(float32Schema.CustomFloat32FuzzFunc())
	float32Fuzzer.Fuzz(&generatedFloat32)
	return generatedFloat32
}
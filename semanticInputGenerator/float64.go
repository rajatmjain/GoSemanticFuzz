package semanticInputGenerator

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
)

func GenerateFloat64(minimum int, maximum int, precision int)(float64){
	var generatedFloat64 float64
	float64Schema := goSemanticFuzz.Float64Schema{Minimum: float64(minimum),Maximum: float64(maximum),Precision: uint(precision)}
	float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&generatedFloat64)
	return generatedFloat64	
}
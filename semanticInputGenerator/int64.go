package semanticInputGenerator

import (
	goSemanticFuzz "GoSemanticFuzz/gosemanticfuzz"
)

func GenerateInt64(minimum int64, maximum int64)(int64){
	var generatedInt64 int64
	int64Schema := goSemanticFuzz.Int64Schema{Minimum:minimum, Maximum:maximum}
	int64Fuzzer := goSemanticFuzz.New().Funcs(int64Schema.CustomInt64FuzzFunc())
	int64Fuzzer.Fuzz(&generatedInt64)
	return (generatedInt64)
}
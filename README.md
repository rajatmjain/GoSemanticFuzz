# GoSemanticFuzz

GoSemanticFuzz is a library for populating go objects with semantically correct random values.

This is useful for testing:

- Do your project's objects really serialize/unserialize correctly in all cases?
- Is there an incorrectly formatted object that will cause your project to panic?

Import with `import "github.com/rajatmjain/GoSemanticFuzz"`

You can use it on single variables:

```go
    // String Fuzzer
	var A string
	unicodeRange := myFuzz.UnicodeRange{First: '0', Last: 'ӿ',MinLength:12,MaxLength: 25}
	stringFuzzer := myFuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	stringFuzzer.Fuzz(&A)
    fmt.Println("Fuzzed string: ",A)

	// Integer64 Fuzzer
	var B int64
	int64Schema := myFuzz.Int64Schema{Minimum: 0,Maximum: 10000}
	int64Fuzzer := myFuzz.New().Funcs(int64Schema.CustomInt64FuzzFunc())
	int64Fuzzer.Fuzz(&B)
	fmt.Println("Fuzzed integer64: ",B)

	// Integer32 Fuzzer
	var C int32
	int32Schema := myFuzz.Int32Schema{Minimum: 0,Maximum: 10000}
	int32Fuzzer := myFuzz.New().Funcs(int32Schema.CustomInt32FuzzFunc())
	int32Fuzzer.Fuzz(&C)
	fmt.Println("Fuzzed integer32: ",C)

	//Float64 Fuzzer
	var D float64
	float64Schema := myFuzz.Float64Schema{Minimum: 0,Maximum: 10000,Precision: 9}
	float64Fuzzer := myFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&D)
	fmt.Println("Fuzzed float64: ",D)

	//Float32 Fuzzer
	var E float32
	float32Schema := myFuzz.Float32Schema{Minimum: 0,Maximum: 10000,Precision: 2}
	float32Fuzzer := myFuzz.New().Funcs(float32Schema.CustomFloat32FuzzFunc())
	float32Fuzzer.Fuzz(&E)
	fmt.Println("Fuzzed float32: ",E)
```

Happy testing!

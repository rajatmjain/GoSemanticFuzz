package semanticInputGenerator

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
)

func GenerateString(first rune,last rune,minLength int,maxLength int)(string){
	var generatedString string
	unicodeRange := goSemanticFuzz.UnicodeRange{First:first, Last:last, MinLength:minLength, MaxLength:maxLength}
	stringFuzzer := goSemanticFuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	stringFuzzer.Fuzz(&generatedString)
    return generatedString
}

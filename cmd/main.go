package main

import (
	testroutines "GoSemanticFuzz/testRoutines"
)

func main(){
	//testroutines.RunGoFuzzTestRoutine()
	// time.Sleep(10*time.Second)
	testroutines.RunGoSemanticFuzzTestRoutine()
	
}

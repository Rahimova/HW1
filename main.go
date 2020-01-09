
package main

import "fmt"

func main() {

	 fueValue := 20.0
	 consumtion := 10.0

	calculationDistanceService  := calculationDistanceService(fueValue, consumtion)
	 
}

func calculationDistanceService(consumtion, fueValue float64 ) float64{

	calculate := (fueValue * 100 )/ consumtion
	return calculate
}

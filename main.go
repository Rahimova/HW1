
package main

import "fmt"

func main() {

	 fueValue := 10.0   // Расстояние в километрах
	 consumtion := 20.0 // Расход бензина за км

	calc  := calculationDistance(fueValue, consumtion)
	fmt.Println(calc)
}

func calculationDistance(consumtion, fueValue float64 ) float64{

	calculate := (fueValue * 100 )/ consumtion
	return calculate
}

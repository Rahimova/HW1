package main

import "testing"

func Test_calculationDistanceService(t *testing.T) {
	 
	tests := []struct {

		fueValue float64
		consumtion float64
		want float64
	}{
		{20.0, 10.0, 50},
	}

	for _, test := range tests{
		got := calculationDistanceService(test.fueValue, test.consumtion)
		if test.want != got {
			t.Error("want: ", test.want, "but got: ", got)
		}
	}
}

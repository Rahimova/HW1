package main

import "testing"

func Test_calculationDistance(t *testing.T) {
	type args struct {
		consumtion float64
		fueValue   float64
	}
	tests := []struct {
		//name string
		fueValue float64
		consumtion float64
		want float64
	}{
		{10.0, 20.0, 200},
	}

	for _, test := range tests{
		got := calculationDistance(test.fueValue, test.consumtion)
		if test.want != got {
			t.Error("want: ", test.want, "but got: ", got)
		}
	}
}
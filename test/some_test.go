package test

import (
	"ConfBackend/util/coord"
	"fmt"
	"testing"
)

func TestLineaerCalc(t *testing.T) {
	x, y, z, err := coord.CalcLinearCoord([][]float64{{2.0, 0.0, 0.0}, {0.0, 0.0, 0.0}}, [][]float64{{2.0}, {2.82842}})
	fmt.Println(x, y, z, err)
	if err != nil {
		return
	}
}

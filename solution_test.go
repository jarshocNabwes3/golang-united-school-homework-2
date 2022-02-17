package square

import (
	"testing"
)

func TestCalcSquare(t *testing.T) {
	var sidesNumArgs = []intCustomType{0, 1, 2, 3, 4, 5,}
	var sidesLenArgs = []float64{1, 1, 1, 1, 1, 1,}
	var expecteds = []float64{1, 0, 0, 1, 1, 0,}

	for i, sidesNum := range sidesNumArgs{
		sideLen, expected := sidesLenArgs[i], expecteds[i]
		var square = CalcSquare(sideLen, sidesNum)
		if expected != square {
			const msg = "Unexpected #%d result of %d sides and %f" +
				" side length:\n\tExpected: %f\n\tGot: %f"
			t.Errorf(msg, i, sidesNum, sideLen, expected, square)
		}
	}

}

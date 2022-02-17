package square

import (
	"testing"

    "gotest.tools/assert"	
)

func TestCalcSquare(t *testing.T) {
	var sidesNumArgs = []intCustomType{0, 1, 2, 3, 4, 5, 3, 3, }
	var sidesLenArgs = []float64{1, 1, 1, 1, 2, 1, 4.4, 15.67}
	var expecteds = []float64{1.5707963267948966, 0, 0, 0.4330127018922193, 4, 0, 8.383125908633367, 106.32579263566237, }

	for i, sidesNum := range sidesNumArgs{
		sideLen, expected := sidesLenArgs[i], expecteds[i]
		var square = CalcSquare(sideLen, sidesNum)
		assert.Equal( t, expected, square )
		if expected != square {
			const msg = "Unexpected #%d result of %d sides and %f" +
				" side length:\n\tExpected: %e\n\tGot: %ee"
			t.Errorf(msg, i, sidesNum, sideLen, expected, square)
		}
	}

}

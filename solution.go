package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

type intCustomType uint8

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const SidesTriangle intCustomType = 3
const SidesSquare intCustomType = 4
const SidesCircle intCustomType = 0

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {

	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen,2) / 2
	case SidesTriangle:
		return sideLen * math.Sqrt(3.0)/4.0
	case SidesSquare:
		return math.Pow(sideLen,2)
	default:
		return 0
	}

}

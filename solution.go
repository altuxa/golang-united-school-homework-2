package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	if sidesNum == 4 {
		return sideLen * sideLen
	} else if sidesNum == 3 {
		a := sideLen * sideLen * 1.73205
		return a / 4
	} else if sidesNum == 0 {
		return sideLen * sideLen * 3.14159265358979323846264338327950288419716939937510582097494459
	} else {
		return 0
	}
	return 0
}

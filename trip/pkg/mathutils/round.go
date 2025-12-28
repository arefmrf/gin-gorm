package mathutils

import "math"

func RoundFloat(val float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(val*pow) / pow
}

package main

func Gps(s int, x []float64) int {
	lastDelta := 0.0
	biggest := 0.0
	for _, i := range x {
		if i-lastDelta > biggest {
			biggest = lastDelta
		}
		lastDelta = i
	}
	return int((3600 * biggest) / float64(s))
}

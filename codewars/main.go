package main

import "fmt"

func main() {
	var x = []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}
	a := Gps(20, x)
	fmt.Println(a)
}

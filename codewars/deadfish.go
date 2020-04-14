package main

func deadfish(data string) []int {
	var output []int = nil
	val := 0

	for _, i := range data {
		switch i {
		case 105:
			val++
		case 100:
			val--
		case 115:
			val *= val
		case 111:
			output = append(output, val)
		}
	}
	return output
}

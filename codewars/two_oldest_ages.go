package main

func TwoOldestAges(ages []int) [2]int {
	oldest := 0
	second_oldest := 0

	for _, i := range ages {
		if i > second_oldest {
			if i > oldest {
				second_oldest = oldest
				oldest = i
			} else {
				second_oldest = i
			}
		}
	}

	return [2]int{second_oldest, oldest}
}

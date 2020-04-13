package main

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	phoneno := "("
	for i, num := range numbers {
		if i == 3 {
			phoneno += ") "
		} else if i == 6 {
			phoneno += "-"
		}
		phoneno += fmt.Sprint(num)
	}

	return phoneno
}
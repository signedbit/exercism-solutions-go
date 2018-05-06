package raindrops

import "strconv"

func Convert(raindrop int) string {
	result := ""
	if raindrop%3 == 0 {
		result += "Pling"
	}
	if raindrop%5 == 0 {
		result += "Plang"
	}
	if raindrop%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		res := strconv.Itoa(raindrop)
		return res
	}
	return result
}

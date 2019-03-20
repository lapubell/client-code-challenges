package main

func makeTriangle(lines int) map[int]string {
	output := make(map[int]string, lines)
	for index := lines; index > 0; index-- {
		output[index] = desiredString(index, lines)
	}
	return output
}

func desiredString(current, total int) string {
	output := ""
	desiredLength := (current * 2) - 1
	for index := 0; index < desiredLength; index++ {
		output = output + outputChar
	}
	return paddedString(output, (total*2)-1)
}

func paddedString(s string, desiredLength int) string {
	difference := desiredLength - len(s)
	if difference == 0 {
		return s
	}
	padding := ""
	for index := 0; index < difference/2; index++ {
		padding = padding + " "
	}
	return padding + s + padding
}

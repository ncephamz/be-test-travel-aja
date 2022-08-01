package main

import "fmt"

func main(){
	// update input value here
	const (
		input1 = "telkom"
		input2 = "tlkom"
	)

	fmt.Println(wordSpelling(input1, input2))
}

func wordSpelling(input1, input2 string) bool {
	if (len(input1) - len(input2) > 2) {
		return false
	}

	charsInput1 := []rune(input1)
	charsInput2 := []rune(input2)

	totalChange := 0

	for idx, char := range charsInput1 {
		if (totalChange > 1) {
			break
		}

		if (char != charsInput2[idx]) {
			if (char == charsInput2[idx + 1]) {
				totalChange += 1
				charsInput2 = insert(charsInput2, char, idx)
			} else {
				totalChange += 1
				charsInput2[idx] = char
			}
		}
	}

	if (totalChange < 1) {
		return true
	} else {
		return false
	}
}

func insert(a []rune, c rune, i int) []rune {
    return append(a[:i], append([]rune{c}, a[i:]...)...)
}

func removeIndex(s []rune, index int) []rune {
	return append(s[:index], s[index+1:]...)
}
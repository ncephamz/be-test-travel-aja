package main

import (
	"fmt"
	"strings"
)

func main(){
	// update input value here
	const (
		input1 = "telkom"
		input2 = "selkom"
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
			if (idx == 0) {
				totalChange += 1
				charsInput2[idx] = char
				continue
			}

			if (idx + 1  == len(charsInput1)){
				if (len(charsInput1) < len(charsInput2)) {
					totalChange += 2
					break
				}
			} 

			if (idx + 1 == len(charsInput2)) {
				totalChange += 1
				charsInput2[idx] = char
				continue
			}

			if (charsInput2[idx] == charsInput2[idx - 1]){
				totalChange += 1
				charsInput2 = remove(charsInput2, idx)
				continue
			}

			if (charsInput2[idx] == charsInput1[strings.Index(input1, string(charsInput2[idx]))]) {
				totalChange += 1
				charsInput2 = insert(charsInput2, char, idx)
				continue
			}
		}
	}

	if (totalChange <= 1) {
		return true
	} else {
		return false
	}
}

func insert(a []rune, c rune, i int) []rune {
    return append(a[:i], append([]rune{c}, a[i:]...)...)
}

func remove(a []rune, i int) []rune {
	return append(a[:i], a[i+1:]...)
}
package php

// SimilarText calculate the similarity between two strings
// returns the number of matching chars in both strings and the similarity percent
// note that there is a little difference between the original php function that
// a multi-byte character is counted as 1 while it would be 2 in php.
// This function is case sensitive.
func SimilarText(first, second string) (count int, percent float32) {
	txt1 := []rune(first)
	len1 := len(txt1)
	txt2 := []rune(second)
	len2 := len(txt2)

	if len1 == 0 || len2 == 0 {
		return 0, 0
	}

	sim := similarChar(txt1, len1, txt2, len2)

	return sim, float32(sim*200) / float32(len1+len2)
}

func similarChar(txt1 []rune, len1 int, txt2 []rune, len2 int) int {
	var sum int

	pos1, pos2, max := similarStr(txt1, len1, txt2, len2)

	if sum = max; sum > 0 {
		if pos1 > 0 && pos2 > 0 {
			sum += similarChar(txt1, pos1, txt2, pos2)
		}
		if (pos1+max < len1) && (pos2+max < len2) {
			sum += similarChar(txt1[(pos1+max):], len1-pos1-max, txt2[(pos2+max):], len2-pos2-max)
		}
	}

	return sum
}

func similarStr(txt1 []rune, len1 int, txt2 []rune, len2 int) (pos1, pos2, max int) {
	var l int
	pos1, pos2, max = 0, 0, 0
	for i := range txt1 {
		for j := range txt2 {
			for l = 0; i+l < len1 && j+l < len2 && txt1[i+l] == txt2[j+l]; l++ {
			}
			if l > max {
				max = l
				pos1 = i
				pos2 = j
			}
		}
	}
	return pos1, pos2, max
}

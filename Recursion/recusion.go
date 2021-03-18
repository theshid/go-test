package Recursion

func Reverse(word string) string {

	if len(word) == 0 {
		return word
	}

	return Reverse(string(word[1:])) + string(word[0])

}

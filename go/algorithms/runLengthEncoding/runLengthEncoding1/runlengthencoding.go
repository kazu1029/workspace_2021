package runlengthencoding1

import (
	"strconv"
)

func RunLengthEncode(str string) string {
	encodedStringCharacters := []byte{}
	currentRunLength := 1

	for i := 1; i < len(str); i++ {
		currentCharacter := str[i]
		previousCharacter := str[i-1]

		if currentCharacter != previousCharacter || currentRunLength == 9 {
			encodedStringCharacters = append(encodedStringCharacters, strconv.Itoa(currentRunLength)[0])
			encodedStringCharacters = append(encodedStringCharacters, previousCharacter)
			currentRunLength = 0
		}

		currentRunLength++
	}

	encodedStringCharacters = append(encodedStringCharacters, strconv.Itoa(currentRunLength)[0])
	encodedStringCharacters = append(encodedStringCharacters, str[len(str)-1])

	return string(encodedStringCharacters)
}

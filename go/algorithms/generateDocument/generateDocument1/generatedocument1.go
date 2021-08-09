package generatedocument1

func GenerateDocument(characters string, document string) bool {
	charsMap := map[rune]int{}
	for _, char := range characters {
		charsMap[char] += 1
	}

	for _, char := range document {
		if found, ok := charsMap[char]; !ok || found == 0 {
			return false
		}
		charsMap[char] -= 1
	}

	return true
}

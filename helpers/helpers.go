package helpers

import "strings"

func GetFirstNames(fullNames []string) []string {
	var firstNames []string
	for _, fullName := range fullNames {
		firstNames = append(firstNames, strings.Fields(fullName)[0])
	}
	return firstNames
}

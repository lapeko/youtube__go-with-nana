package helpers

func GetFirstNames(orders []map[string]string) []string {
	var firstNames []string
	for _, order := range orders {
		firstNames = append(firstNames, order["firstName"])
	}
	return firstNames
}

package helpers

import "booking-app/models"

func GetFirstNames(orders []models.Order) []string {
	var firstNames []string
	for _, order := range orders {
		firstNames = append(firstNames, order.FirstName)
	}
	return firstNames
}

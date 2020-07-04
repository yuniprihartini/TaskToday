package utils

import "ujian/master/models"

func ValidationServices(input *models.Services) bool {
	var result bool
	result = true
	switch {
	case input.VehicleType == "":
		result = false
	case input.ServiceType == "":
		result = false
	}
	return result
}
func ValidationProgress(input *models.Progress) bool {
	var result bool
	result = true
	switch {
	case input.CustomerId == "":
		result = false
	case input.ServiceId == "":
		result = false
	}
	return result
}
func ValidationCustomers(input *models.Customers) bool {
	var result bool
	result = true
	switch {
	case input.CustomerName == "":
		result = false
	case input.BookingDate == "":
		result = false
	case input.NumberPlat == "":
		result = false
	}
	return result
}

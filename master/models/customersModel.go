package models

type Customers struct {
	CustomerId   string `json:"customerid"`
	CustomerName string `json:"customername"`
	BookingDate  string `json:"booking_date"`
	BookingTime  string `json:"booking_time"`
	NumberPlat   string `json:"plat_number"`
	ServiceID    string `json:"service_id"`
}

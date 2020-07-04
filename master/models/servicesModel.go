package models

type Services struct {
	ServiceId    string `json:"service_id"`
	ServiceType  string `json:"service_type"`
	VehicleType  string `json:"vehicle_type"`
	ServicePrice int    `json:"service_price"`
}

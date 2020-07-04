package models

type Progress struct {
	ProgressId    string `json:"progress_id"`
	ServiceId     string `json:"service_id"`
	CustomerId    string `json:"customer_id"`
	ServiceStatus string `json:"service_status"`
}

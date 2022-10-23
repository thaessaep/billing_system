package model

type ReserveBills struct {
	Success   *bool
	OrderId   int `json:"order_id"`
	ServiceId int `json:"service_id"`
	Cost      int `json:"cost"`
	Datetime  int64
	User
}

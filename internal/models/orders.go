package models

import (
	"net/http"
	"encoding/json"
)

type OrderID struct {
	OrderID uint64 `json:"order_id"`
}

func (o *OrderID) DataValidation(r *http.Request) bool {
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
        return false
    }

	if o.OrderID < 1 {
		return false 
	}

	return true
}

type Order struct {
	UserID uint64 `json:"user_id"`
	OrderID uint64 `json:"order_id"`
	ServiceID uint64 `json:"service_id"`
	ServiceName string `json:"service_name"`
	Cost float64 `json:"cost"`
}

func (o *Order) DataValidation(r *http.Request) bool {
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
        return false
    }

	if o.UserID < 1 || o.OrderID < 1 || o.ServiceID < 1 {
		return false 
	}
	
	if o.Cost < 0 {
		return false 
	}

	if o.ServiceName == "" {
		return false 
	}

	return true
}
package models

import (
	"net/http"
	"encoding/json"
)


type UserID struct {
	UserID uint64 `json:"user_id"`
}

func (u *UserID) DataValidation(r *http.Request) bool {
	err := json.NewDecoder(r.Body).Decode(&u) 
	if err != nil {
        return false
    }

	if u.UserID < 1 {
		return false
	}

	return true
}

type DepBalance struct {
	UserID uint64 `json:"user_id"`
	Deposit float64 `json:"deposit"`
	Comment string `json:"comment"`
}

func (d *DepBalance) DataValidation(r *http.Request) bool {
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
        return false
    }

	if d.UserID < 1 || d.Deposit < 0 {
		return false 
	}

	if d.Comment == "" {
		return false 
	}
	
	return true
}

type Balance struct {
	Balance float64 `json:"balance"`
}
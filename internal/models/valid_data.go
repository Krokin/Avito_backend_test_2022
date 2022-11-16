package models

import (
	"net/http"
)

type Approval interface {
	DataValidation(r *http.Request) (bool)
}

func ValidData(a Approval, r *http.Request) bool {
	return a.DataValidation(r)
}
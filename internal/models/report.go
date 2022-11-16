package models

import (
	"net/http"
	"strings"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"
)


type DownloadURL struct {
	URL string `json:"url"`
}

func (d *DownloadURL) FormatURL(r *ReportDate) {
	d.URL = fmt.Sprintf(`http://127.0.0.1:4000/api/sale/sum_report/download/%d_%d`, r.Month, r.Year)
}

type ReportDate struct {
	Year uint64 `json:"year"`
	Month uint64 `json:"month"`
}

func (d *ReportDate) DataValidation(r *http.Request) bool {
	if d.Month < 1 || d.Month > 12 {
		return false
	}

	if d.Year < 0 {
		return false
	}
	
	return true
}

func (d *ReportDate) ParseDate(r *http.Request) bool {
	params := mux.Vars(r)
    f, err := params["date"]
	if err == false {
		return false
	}

	date := strings.Split(f, "_")
	if len(date) != 2 {
		return false
	}

	year, erro := strconv.ParseUint(date[1], 10, 64)
	if erro != nil {
		return false
	}
	d.Year = year

	month, erro := strconv.ParseUint(date[0], 10, 64)
	if erro != nil {
		return false
	}
	d.Month = month

	return true
}
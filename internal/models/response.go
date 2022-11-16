package models

import (
	"net/http"
	"encoding/json"
	"encoding/csv"
	"fmt"
)

type Response interface {
	RespWrite(w http.ResponseWriter) error
}

func ResponseWrite(a Response, w http.ResponseWriter) error {
	err := a.RespWrite(w)

	return err
}

func (b *Balance) RespWrite(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(b)

	return err
}

func (d *DepBalance) RespWrite(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	r := &Error{
		ErrorCode: 0, 
		ErrorMessage: "none error message",
	}
	err := json.NewEncoder(w).Encode(r)

	return err
}

func (o *Order) RespWrite(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	r := &Error{
		ErrorCode: 0, 
		ErrorMessage: "none error message",
	}
	err := json.NewEncoder(w).Encode(r)

	return err
}

func (d *DownloadURL) RespWrite(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(d)

	return err
}

func (r *ReportDate) RespWrite(w http.ResponseWriter, reportList [][]string) error {
	w.Header().Set("Content-Type", "text/csv")

	fileName := fmt.Sprintf("attachment;filename=report_%d_%d.csv", r.Year, r.Month)
	w.Header().Set("Content-Disposition", fileName)

    w.WriteHeader(http.StatusOK)

	wr := csv.NewWriter(w)
	err := wr.WriteAll(reportList)

	return err
}

func (h *HistoryReq) RespWrite(w http.ResponseWriter, listTransaction []Transaction) error {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(TransactionResponse{Transaction_list: listTransaction})
	
	return err
}
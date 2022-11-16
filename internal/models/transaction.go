package models

import (
	"net/http"
	"strconv"
	"fmt"
)

type HistoryReq struct {
	UserID int
	PageNo int
	PageSize int
	Sort string
	Ascending bool
}

func (h *HistoryReq) DataValidation(r *http.Request) bool {
	err := true
	var ok error

	h.UserID, ok = strconv.Atoi(r.URL.Query().Get("user_id"))
	if ok != nil || h.UserID < 1 {
		err = false
	}

	h.PageNo, ok = strconv.Atoi(r.URL.Query().Get("page_no"))
	if ok != nil || h.PageNo < 1 {
		err = false
	}

	h.PageSize, ok = strconv.Atoi(r.URL.Query().Get("page_size"))
	if ok != nil || h.PageSize > 100 || h.PageSize < 1 {
		err = false
	}

	h.Sort = r.URL.Query().Get("sort")
	if h.Sort != "date" && h.Sort != "cost" {
		err = false
	}
	
	h.Ascending, ok = strconv.ParseBool(r.URL.Query().Get("ascending"))
	if ok != nil {
		err = false
	}

	return err
}

func (h *HistoryReq) FormatQuery() string {
	var sort = "date"
	if h.Sort != "date" {
		sort = "sum_transaction"
	}

	if h.Ascending == true {
		sort += " ASC"
	} else {
		sort += " DESC"
	}
	
	query := fmt.Sprintf(`SELECT id, type_transaction, sum_transaction, comment, date FROM user_transaction WHERE user_id = ? ORDER BY %s LIMIT ?`, sort)
	if h.PageNo != 1 {
		offset := (h.PageNo-1) * h.PageSize
		query += fmt.Sprintf(", %d", offset)
	}
	
	return query
}

type Transaction struct {
	TransactionID int `json:"transaction_id"`
	TypeTransaction string `json:"type_transaction"`
	Cost float64 `json:"cost"`
	Comment string `json:"comment"`
	Date string  `json:"date"`
}

type TransactionResponse struct {
	Transaction_list []Transaction `json:"list_transactions"`
}

package database

import (
	m "Avito_tech_test_2022/internal/models"
)


func (d *DB) HistoryTransaction(h *m.HistoryReq) ([]m.Transaction, error) {
	q := h.FormatQuery()
	
	rows, err := d.DB.Query(q, h.UserID, h.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listTransaction := []m.Transaction{}
	for rows.Next() {
		t := m.Transaction{}
		if err := rows.Scan(&t.TransactionID, &t.TypeTransaction, &t.Cost, &t.Comment, &t.Date); err != nil {
			return nil, err
		}
		listTransaction = append(listTransaction, t)
	}

	return listTransaction, nil
}
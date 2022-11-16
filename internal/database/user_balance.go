package database

import (
	m "Avito_tech_test_2022/internal/models"
)

func (d *DB) GetBalance(uID *m.UserID, uB *m.Balance ) error {
	err := d.DB.QueryRow(`SELECT balance FROM user_balance WHERE user_id = ?`, uID.UserID).Scan(&uB.Balance)
	if err != nil {
		return err
	}
	
	return nil
}

func (d *DB) DepositBalance(dR *m.DepBalance) error {
	_, err := d.DB.Exec(`UPDATE user_balance SET balance = balance+? WHERE user_id = ?`, dR.Deposit, dR.UserID)
	if err != nil {
		return err
	}

	_, err = d.DB.Exec(`INSERT INTO user_transaction (user_id, type_transaction, sum_transaction, comment) VALUES (?, 'deposit', ?, ?)`, 
	dR.UserID, dR.Deposit, dR.Comment)
	if err != nil {
		return err
	}

	return nil
}
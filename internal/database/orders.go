package database

import (
	"fmt"

	m "Avito_tech_test_2022/internal/models"
)

func (d *DB) Reserve(o *m.Order) error {
	tx, err := d.DB.Begin()
    if err != nil {
		return err
    }
    defer tx.Rollback()

    _, err = tx.Exec(`UPDATE user_balance SET balance = balance-? WHERE user_id = ?`, o.Cost, o.UserID )
    if err != nil {
		return err
    }

    _, err = tx.Exec(`INSERT INTO orders_balance (order_id, user_id, service_id, service_name, cost) VALUES (?, ?, ?, ?, ?)`, 
		o.OrderID, o.UserID, o.ServiceID, o.ServiceName, o.Cost)
    if err != nil {
		return err
    }

	comment := fmt.Sprintf("Оплата услуги: %d.%s", o.ServiceID, o.ServiceName)
	_, err = tx.Exec(`INSERT INTO user_transaction (user_id, type_transaction, sum_transaction, comment) VALUES (?, 'withdraw', ?, ?)`, 
	o.UserID, o.Cost, comment)
	if err != nil {
		return err
	}

    if err = tx.Commit(); err != nil {
		return err
    }
	
	return nil
}

func (d *DB) ReserveOut(oID *m.OrderID, o *m.Order) error {
	err := d.DB.QueryRow(`SELECT * FROM orders_balance WHERE order_id = ?`, oID.OrderID).
		Scan(&o.OrderID, &o.UserID, &o.ServiceID, &o.ServiceName, &o.Cost)
	if err != nil {
		return err
	}

    tx, err := d.DB.Begin()
    if err != nil {
		return err
    }

    defer tx.Rollback()

    _, err = tx.Exec(`UPDATE user_balance SET balance = balance+? WHERE user_id = ?`, o.Cost, o.UserID )
    if err != nil {
		return err
    }

    _, err = tx.Exec(`DELETE FROM orders_balance WHERE order_id = ?`, o.OrderID)
    if err != nil {
		return err
    }

	comment := fmt.Sprintf("Возврат оплаты за отмену услуги: %d.%s", o.ServiceID, o.ServiceName)
	_, err = tx.Exec(`INSERT INTO user_transaction (user_id, type_transaction, sum_transaction, comment) VALUES (?, 'withdraw', ?, ?)`, 
	o.UserID, o.Cost, comment)
	if err != nil {
		return err
	}

    if err = tx.Commit(); err != nil {
		return err
    }

	return nil
}

func (d *DB) Revenue(oID *m.OrderID, o *m.Order) error {
	err := d.DB.QueryRow(`SELECT * FROM orders WHERE order_id = ?`, o.OrderID).
		Scan(&o.OrderID, &o.UserID, &o.ServiceID, &o.Cost)
	if err != nil {
		return err
	}

    tx, err := d.DB.Begin()
    if err != nil {
		return err
    }

    defer tx.Rollback()

    _, err = tx.Exec(`INSERT INTO orders_balance (order_id, user_id, service_id, service_name, cost) VALUES (?, ?, ?, ?, ?)`, 
	o.OrderID, o.UserID, o.ServiceID, o.ServiceName, o.Cost)
    if err != nil {
		return err
    }

    _, err = tx.Exec(`DELETE FROM orders WHERE order_id = ?`, o.OrderID)
    if err != nil {
		return err
    }

    if err = tx.Commit(); err != nil {
		return err
    }
	
	return nil
}

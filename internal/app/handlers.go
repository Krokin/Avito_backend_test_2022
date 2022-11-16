package app

import (
	"net/http"

	m "Avito_tech_test_2022/internal/models"
)

func (app *Application) GetBalance(w http.ResponseWriter, r *http.Request) {
	userID := &m.UserID{}
	erro := m.ValidData(userID, r)
	if erro == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	userBalance := &m.Balance{}
	err := app.DB.GetBalance(userID, userBalance)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.ResponseWrite(userBalance, w)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) DepositBalance(w http.ResponseWriter, r *http.Request) {
	deposRequest := &m.DepBalance{}
	erro := m.ValidData(deposRequest, r)
	if erro == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.DB.DepositBalance(deposRequest)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.ResponseWrite(deposRequest, w)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) Reserve(w http.ResponseWriter, r *http.Request) {
	order := &m.Order{}
	erro := m.ValidData(order, r)
	if erro == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}

    err := app.DB.Reserve(order)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.ResponseWrite(order, w)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) ReserveOut(w http.ResponseWriter, r *http.Request) {
	orderID := &m.OrderID{}
	erro := m.ValidData(orderID, r)
	if erro == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	order := &m.Order{}
	err := app.DB.ReserveOut(orderID, order)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.ResponseWrite(order, w)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) Revenue(w http.ResponseWriter, r *http.Request) {
	orderReq := &m.OrderID{}
	erro := m.ValidData(orderReq, r)
	if erro == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	orderData := &m.Order{}
	err := app.DB.Revenue(orderReq, orderData)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.ResponseWrite(orderData, w)
	if err != nil {
		app.serverError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) SumReport(w http.ResponseWriter, r *http.Request) {
	reportDate := &m.ReportDate{}
	err := m.ValidData(reportDate, r)
	if err == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	url := &m.DownloadURL{}
	url.FormatURL(reportDate)
	
	error := m.ResponseWrite(url, w)
	if error != nil {
		app.serverError(w, error.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) DownloadReport(w http.ResponseWriter, r *http.Request) {
	date := &m.ReportDate{}
	if err := date.ParseDate(r); err == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := m.ValidData(date, r)
	if err == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	reportList, error := app.DB.DownloadReport(date)
	if error != nil {
		app.serverError(w, error.Error(), http.StatusInternalServerError)
		return
	}

	error = date.RespWrite(w, reportList)
	if error != nil {
		app.serverError(w, error.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) HistoryTransaction(w http.ResponseWriter, r *http.Request) {
	historyReq := &m.HistoryReq{}
	err := m.ValidData(historyReq, r)
	if err == false {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	
	listTransaction, error := app.DB.HistoryTransaction(historyReq)
	if error != nil {
		app.serverError(w, error.Error(), http.StatusInternalServerError)
		return
	}

	error = historyReq.RespWrite(w, listTransaction)
	if error != nil {
		app.serverError(w, error.Error(), http.StatusInternalServerError)
		return
	}
}


package app

import (
	"net/http"

	m "Avito_tech_test_2022/internal/models"
)

// @Summary GetBalance
// @Tags user_balance
// @Description get user balance
// @ID get-balance
// @Accept json
// @Produce json
// @Param input body models.UserID true "user id"
// @Success 200 {object} models.Balance
// @Failure 400
// @Failure 500
// @Router /api/ub/get_balance [post]
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

// @Summary DepositBalance
// @Tags user_balance
// @Description deposit user balance
// @ID deposit-balance
// @Accept json
// @Produce json
// @Param input body models.DepBalance true "user id, sum deposit, comment"
// @Success 200 {object} models.Error
// @Failure 400
// @Failure 500
// @Router /api/ub/deposit_balance [post]
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

// @Summary Reserve
// @Tags sale
// @Description reserving funds from the user's balance
// @ID reserve
// @Accept json
// @Produce json
// @Param input body models.Order true "data order"
// @Success 200 {object} models.Error
// @Failure 400
// @Failure 500
// @Router /api/sale/reserve [post]
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

// @Summary ReserveOut
// @Tags sale
// @Description return of reserved funds to the user
// @ID reserveout
// @Accept json
// @Produce json
// @Param input body models.OrderID true "order id"
// @Success 200 {object} models.Error
// @Failure 400
// @Failure 500
// @Router /api/sale/reserve_out [post]
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

// @Summary Revenue
// @Tags sale
// @Description profit recognition
// @ID revenue
// @Accept json
// @Produce json
// @Param input body models.OrderID true "order id"
// @Success 200 {object} models.Error
// @Failure 400
// @Failure 500
// @Router /api/sale/revenue [post]
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

// @Summary SumReport
// @Tags sale
// @Description getting a link to a file for a report 
// @ID sum_report
// @Accept json
// @Produce json
// @Param input body models.ReportDate true "date for report"
// @Success 200 {object} models.DownloadURL
// @Failure 400
// @Failure 500
// @Router /api/sale/sum_report [post]
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

// @Summary SumReport_Download
// @Tags sale
// @Description download report 
// @ID sum_report_download
// @Accept json
// @Produce mpfd
// @Param date query string true "date for report" Format(month_year)
// @Success 200 {array} []string
// @Failure 400
// @Failure 500
// @Router /api/sale/sum_report/download/{date} [get]
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


// @Summary HistoryTransaction
// @Tags sale
// @Description history user transaction
// @ID history_transaction
// @Accept json
// @Produce json
// @Param user_id query int true "user id"
// @Param page_no query int true "page number"
// @Param page_size query int true "page size (max 100)"
// @Param sort query string true "cost or date" format(cost/date)
// @Param ascending query bool true "ASC = true, DESC = false"
// @Success 200 {array} models.Transaction
// @Failure 400
// @Failure 500
// @Router /api/ub/history_transactions [get]
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


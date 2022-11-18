package app

import (
	"github.com/gorilla/mux"

    "github.com/swaggo/http-swagger"
	// _ "github.com/swaggo/http-swagger/example/gorilla/docs"
    _ "Avito_tech_test_2022/docs"
)

func (app *Application) Routes() *mux.Router {
    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/api/ub/get_balance", app.GetBalance).
        Methods("POST")
    r.HandleFunc("/api/ub/deposit_balance", app.DepositBalance).
        Methods("POST")
    r.HandleFunc("/api/ub/history_transactions", app.HistoryTransaction).
        Methods("GET")
    r.HandleFunc("/api/sale/reserve", app.Reserve).
        Methods("POST")
    r.HandleFunc("/api/sale/reserve_out", app.ReserveOut).
        Methods("POST")
    r.HandleFunc("/api/sale/revenue", app.Revenue).
        Methods("POST")
    r.HandleFunc("/api/sale/sum_report", app.SumReport).
        Methods("POST")
    r.HandleFunc("/api/sale/sum_report/download/{date}", app.DownloadReport).
        Methods("GET")
    r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
    return r
}

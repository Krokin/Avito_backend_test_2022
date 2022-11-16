package app

import (
	"log"
	"net/http"
	"fmt"
	"runtime/debug"

	d "Avito_tech_test_2022/internal/database"
)

type Application struct {
    ErrorLog      *log.Logger
    InfoLog       *log.Logger
    DB            *d.DB
}

func (a *Application) clientError(w http.ResponseWriter, err int) {
	http.Error(w, http.StatusText(err), err)
}

func (a *Application) serverError(w http.ResponseWriter, comment string, err int) {
	trace := fmt.Sprintf("%s\n%s", comment, debug.Stack())
	a.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(err), err)
}

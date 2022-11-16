package main

import (
	"log"
	"net/http"
	"os"
	"time"

    app "Avito_tech_test_2022/internal/app"
    d "Avito_tech_test_2022/internal/database"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    sqlbase, err := d.OpenDB("root:199810@/avito_backend")
    if err != nil {
        errorLog.Fatal(err)
    }
    defer sqlbase.Close()

    app := &app.Application{
        ErrorLog:      errorLog,
        InfoLog:       infoLog,
        DB:            &d.DB{DB: sqlbase},
    }

    srv := &http.Server{
        Addr: "0.0.0.0:8080",
        WriteTimeout: time.Second * 15,
        ReadTimeout: time.Second * 15,
        Handler: app.Routes(),
    }
    app.InfoLog.Printf("Starting the server http://127.0.0.1%s", ":4000")
    app.ErrorLog.Fatal(srv.ListenAndServe())
}


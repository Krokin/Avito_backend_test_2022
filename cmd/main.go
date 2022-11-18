package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	app "Avito_tech_test_2022/internal/app"
	"Avito_tech_test_2022/internal/config"
	d "Avito_tech_test_2022/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    cfg, err := config.GetConfig()
	if err != nil {
		errorLog.Fatal(err)
	}

    sqlbase, err := d.OpenDB(cfg.MySQL)
    if err != nil {
        errorLog.Print(err)
    }
    defer sqlbase.Close()

    app := &app.Application{
        ErrorLog:      errorLog,
        InfoLog:       infoLog,
        DB:            &d.DB{DB: sqlbase},
    }

    srv := &http.Server{
        Addr:           fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
        WriteTimeout:   time.Second * 15,
        ReadTimeout:    time.Second * 15,
        Handler:        app.Routes(),
    }

    app.InfoLog.Printf("Starting the server http://localhost:%s", cfg.Port)
    app.ErrorLog.Fatal(srv.ListenAndServe())
}


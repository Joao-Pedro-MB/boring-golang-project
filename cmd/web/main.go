package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Joao-Pedro-MB/boring-golang-project/internal/models"
	_ "github.com/go-sql-driver/mysql" // New import
	"github.com/spf13/viper"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	messages *models.MessageModel
}

// this whole function is just a exageration to practice th usage of a .env file
func viperEnvVariable(key string) string {

	viper.SetConfigFile("../../dev.env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func main() {
	default_addr := viperEnvVariable("DEFAULT_ADDR")
	addr := flag.String("addr", default_addr, "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := flag.String("dsn", "user:password@/boringProject?parseTime=true", "MySQL data source name")
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		messages: &models.MessageModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

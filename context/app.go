package context

import (
	"cabal-api/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

// App struct
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

var setting *config.Mysql

func init() {
	setting = &config.Setting.App.Mysql
}

// Run do start server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// InitializeRoute is create app constance
func (a *App) InitializeRoute() {
	a.Router = mux.NewRouter()
	// topup
	a.Router.HandleFunc("/topup", a.CashHandler).Methods(http.MethodPost)
	// topup callback
	a.Router.HandleFunc("/topup/callback", a.CashCallBackHandler).Methods(http.MethodPost)
	// items
	a.Router.HandleFunc("/items", a.ItemMultipleHandler).Methods(http.MethodPost)
	// account info
	a.Router.HandleFunc("/users/{user_id}", a.InfoHandler).Methods(http.MethodGet)
	// migrate info
	a.Router.HandleFunc("/migrate", a.MigrateHandler).Methods(http.MethodPost)
	// login
	a.Router.HandleFunc("/login", a.LoginHandler).Methods(http.MethodPost)
	// Characters List
	a.Router.HandleFunc("/characters/{user_id}", a.CharacterHandler).Methods(http.MethodGet)

	// ResetSubPassword List
	a.Router.HandleFunc("/sub-password/{user_id}", a.ResetSubPasswordHandler).Methods(http.MethodPut)
}

// InitializeDatabase is create database instance
func (a *App) InitializeDatabase() {
	var err error
	strConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.Username, setting.Password, setting.Host, setting.Database)
	a.DB, err = sql.Open("mysql", strConnection)
	if err != nil {
		panic(err)
	}
	a.DB.SetConnMaxLifetime(0)
}

// CheckAndRetryConnection do checking connection and retry them
func (a *App) CheckAndRetryConnection() error {
	var err error
	err = a.DB.Ping()
	if err != nil {
		strConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.Username, setting.Password, setting.Host, setting.Database)
		a.DB, err = sql.Open("mysql", strConnection)
		if err != nil {
			return fmt.Errorf("error RetryConnection %s", err.Error())
		}
	}

	return err
}

// InitializeDatabaseTest is create database instance
func (a *App) InitializeDatabaseTest() sqlmock.Sqlmock {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	a.DB = db
	return mock
}

// Exec Database retry and handler Error.
func (a *App) Exec(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	res, err := stmt.Exec(args...)
	if err != nil {
		for strings.Contains(strings.ToLower(err.Error()), "deadlock") {
			time.Sleep(time.Millisecond * 10)
			res, err = stmt.Exec(args...)
			if err == nil {
				break
			}
		}
	}

	return res, err
}

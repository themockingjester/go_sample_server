package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

// DBConfig represents an SQL database's configuration.
type DBConfig struct {
	Type           string        `mapstructure:"type"`
	Endpoint       string        `mapstructure:"dsn"`
	MaxIdleConns   int           `mapstructure:"max_idle"`
	MaxActiveConns int           `mapstructure:"max_active"`
	ConnectTimeout time.Duration `mapstructure:"connect_timeout"`
}

var DbConnection *sqlx.DB

func main() {

	//creating new database connection
	cfg := DBConfig{
		Type:           "mysql",
		Endpoint:       "root:Taken123@@tcp(localhost:3306)/go_sample_server",
		MaxIdleConns:   10,
		MaxActiveConns: 100,
		ConnectTimeout: 10,
	}
	connection, dbErr := ConnectDB(cfg)
	if dbErr != nil {
		fmt.Printf("Unable to connect to database because %v", dbErr)
		os.Exit(1)
	}
	DbConnection = connection
	// Bind the server HTTP endpoints.
	r := chi.NewRouter()

	r.Post("/postRequest/{source}", PostRequest)
	r.HandleFunc("/about", AboutPage)
	r.Get("/getRequest", GetRequest)

	r.Post("/addUser", AddUser)
	err := http.ListenAndServe(":9098", r)

	// These below print statements are not working as expected
	if err == nil {
		fmt.Println("Successfully started server")
	} else {
		fmt.Println("An Error occoured  %v", err)
	}
}

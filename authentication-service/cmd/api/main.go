package main

//This is going to be a web service/ api , need to listen on specific port
import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	//Todo connect to DB
	conn := connectToDb()
	if conn == nil {
		log.Panic("Cant connect to PostGres!")
	}

	//connect to database postgres

	//setup config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	//setup webserver
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}

// This function takes connection string and returns pointer to SqlDB and error
func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil

}

// we also need to add postgres on docker compose file, we need to make sure its available before we return this connection, as this service might startup before db does
// So another function
// we take dsn from environment and then add to docker compose
func connectToDb() *sql.DB {

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}

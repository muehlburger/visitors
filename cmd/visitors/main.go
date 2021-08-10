package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
	"visitors.it-zt.at/infrastructure/repository"
	"visitors.it-zt.at/usecase/measurement"
)

var (
	service *measurement.Service
)

const version string = "1.0.0"

type config struct {
	env string
	db  struct {
		file        string
		maxOpenCons int
		maxIdleCons int
		maxIdleTime string
	}
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.file, "db-file", os.Getenv("SQLITE_DB_FILE"), "SQLite database file. E.g. visitors.db")

	flag.Parse()

	app := &application{
		config: cfg,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	db, err := openDB(cfg)
	if err != nil {
		app.logger.Fatalf("Could not create database file: %v", err)
	}

	app.initDB(db)
	defer db.Close()
	app.logger.Printf("database connection pool established")

	repo := repository.NewStorage(db)
	service = measurement.NewService(repo)

	app.logger.Printf("crawler started on %s", cfg.env)
	doEvery(5*time.Second, crawl)
}

func (app *application) initDB(db *sql.DB) {
	createTableStmt := `CREATE TABLE IF NOT EXISTS visits (
				"id" STRING NOT NULL PRIMARY KEY,
				"quantity" INTEGER,
				"created_at" DATETIME
			);`

	log.Println("create visits table ...")
	stmt, err := db.Prepare(createTableStmt)
	if err != nil {
		app.logger.Fatalf("could not create table: %v", err)
	}
	stmt.Exec()
	app.logger.Println("visits table created")

	createIndexStmt := "CREATE UNIQUE INDEX IF NOT EXISTS idx_visits_id ON visits(id);"
	stmt, err = db.Prepare(createIndexStmt)
	if err != nil {
		app.logger.Fatalf("could not create table: %v", err)
	}
	stmt.Exec()
	log.Println("visits index created")

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.db.file)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func crawl(t time.Time) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.gleisdorf.at"),
	)
	c.OnHTML(`div[class="schwimmbadzaehlerBox"]`, func(e *colly.HTMLElement) {
		v := e.ChildText("strong:first-child")
		service.CreateReading(visitors(v))
	})

	c.Visit("https://www.gleisdorf.at/wellenbad_314.htm")

}

// visitors returns the parsed value of current visitors
func visitors(text string) int {
	v := strings.Split(text, "/")
	visitors, err := strconv.Atoi(v[0])
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	return visitors
}

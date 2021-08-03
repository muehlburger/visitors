package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
	"visitors.it-zt.at/infrastructure/repository"
	"visitors.it-zt.at/usecase/measurement"
)

var (
	once    sync.Once
	service *measurement.Service
)

func main() {
	once.Do(func() {
		db, err := sql.Open("sqlite3", "visitors.db")
		if err != nil {
			log.Fatalf("Could not create database file: %v", err)
		}

		createTableStmt := `CREATE TABLE IF NOT EXISTS visits (
				"id" STRING NOT NULL PRIMARY KEY,
				"quantity" INTEGER,
				"created_at" DATETIME
			);`

		log.Println("Create visits table ...")
		stmt, err := db.Prepare(createTableStmt)
		if err != nil {
			log.Fatalf("Could not create table: %v", err)
		}
		stmt.Exec()
		log.Println("Visits table created")

		createIndexStmt := "CREATE UNIQUE INDEX IF NOT EXISTS idx_visits_id ON visits(id);"
		stmt, err = db.Prepare(createIndexStmt)
		if err != nil {
			log.Fatalf("Could not create table: %v", err)
		}
		stmt.Exec()
		log.Println("Visits Index created")

		repo := repository.NewMeasurementSQLite(db)
		service = measurement.NewService(repo)
	})

	for i := 0; i < 1024; i++ {

	}

	doEvery(5*time.Second, crawl)
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
		currentVisitors := visitors(v)
		fmt.Printf("%v;%v;%d;%.0f\n", t, t.Unix(), currentVisitors, fraction(currentVisitors))
		service.CreateMeasurement(currentVisitors)
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

// fraction calculates the fraction of visitors out of 1500
func fraction(visitors int) float64 {
	return float64(visitors) / 1500 * 100
}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
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

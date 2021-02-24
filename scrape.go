package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type data = struct {
	Name         string
	ClockTimes   string
	Type         string
	Rated        string
	TotalTime    string
	Winner       string
	TotalPlayers string
	Link         string
	Date         string
}

func createCSV(fileName string, header []string) csv.Writer {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fileName, err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write(header)

	return *writer
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	fileName := "Chess_Tournaments.csv"
	ChessTournaments := createCSV(fileName, []string{"Name", "Clock Times", "Type", "Rated", "Total Time", "Winner", "Total Players", "Link", "Date"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#main-wrap > main > div > div > table > tbody > tr", func(e *colly.HTMLElement) {
		Name := e.ChildText("td.header > a > span.name")
		items := strings.Split(e.ChildText("td.header > a > span:nth-child(2)"), " • ")
		ClockTimes := items[0]
		Type := items[1]
		Rated := items[2]
		TotalTime := items[3]
		Winner := e.ChildText("td.winner")
		TotalPlayers := e.ChildText("td.text")
		id := e.ChildAttr("td.header > a", "href")
		Link := "https://lichess.org" + id
		Date := e.ChildAttr("td:nth-child(3) > time", "datetime")

		scrappedData := data{Name, ClockTimes, Type, Rated, TotalTime, Winner, TotalPlayers, Link, Date}

		ChessTournaments.Write([]string{Name, ClockTimes, Type, Rated, TotalTime, Winner, TotalPlayers, Link, Date})

		// fmt.Printf("Found: %q\n", e.Text)
		fmt.Printf("Found: %q\n", scrappedData)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on
	c.Visit("https://lichess.org/@/cbarnes2000/tournaments/created")

	log.Printf("Scraping finished, check file %q for results\n", fileName)
}

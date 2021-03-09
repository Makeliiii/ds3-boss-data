package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"log"
	"github.com/gocolly/colly"
	fatih "github.com/fatih/structs"
	structs "bossScraper/structs"
)

func main() {
	// create a file to store boss data
	file, err := os.Create("data/bosses.csv")

	// check file creation error
	if err != nil {
		log.Fatal(err)
	}

	// wait for func to return then close file
	defer file.Close()

	// create a writer, wait for func to return and write
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// create a collector
	c := colly.NewCollector(
		colly.AllowedDomains("darksouls3.wiki.fextralife.com"),
	)

	// log some stuff when visiting url
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// search for tbody element
	c.OnHTML("tbody", func(e *colly.HTMLElement) {

		// loop through rows and create a boss for each
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			newBoss := structs.Boss {
				Boss: row.ChildText("td:nth-child(1)"),
				Location: row.ChildText("td:nth-child(2)"),
				NPCSummoning: row.ChildText("td:nth-child(3)"),
				Weakness: row.ChildText("td:nth-child(4)"),
				Resistance: row.ChildText("td:nth-child(5)"),
				Immunity: row.ChildText("td:nth-child(6)"),
				Parryable: row.ChildText("td:nth-child(7)"),
				Optional: row.ChildText("td:nth-child(8)"),
			}


			boss := make([]string, 0)
			for _, key := range fatih.Values(newBoss) {
				boss = append(boss, key.(string))
			}

			// write to csv and check for err
			err := writer.Write(boss)
			if err != nil {
				log.Fatal(err)
			}
		})
	})

	// url to scrape
	c.Visit("https://darksouls3.wiki.fextralife.com/Bosses")
}
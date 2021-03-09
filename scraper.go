package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"github.com/gocolly/colly"
	structs "bossScraper/structs"
)

func main() {
	// create a collector
	c := colly.NewCollector(
		colly.AllowedDomains("darksouls3.wiki.fextralife.com"),
	)

	// log some stuff when visiting url
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// create a slice to store bosses
	bosses := []structs.Boss{}

	// search for tbody element
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		// loop through rows and create a boss for each
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			boss := structs.Boss {
				Boss: row.ChildText("td:nth-child(1)"),
				Location: row.ChildText("td:nth-child(2)"),
				NPCSummoning: row.ChildText("td:nth-child(3)"),
				Weakness: row.ChildText("td:nth-child(4)"),
				Resistance: row.ChildText("td:nth-child(5)"),
				Immunity: row.ChildText("td:nth-child(6)"),
				Parryable: row.ChildText("td:nth-child(7)"),
				Optional: row.ChildText("td:nth-child(8)"),
			}

			// append boss to bosses
			bosses = append(bosses, boss)
		})
		fmt.Println(bosses)
	})

	// url to scrape
	c.Visit("https://darksouls3.wiki.fextralife.com/Bosses")
}
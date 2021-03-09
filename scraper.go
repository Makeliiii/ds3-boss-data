package main

import (
	"fmt"
	"github.com/gocolly/colly"
	structs "bossScraper/structs"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("darksouls3.wiki.fextralife.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	bosses := []structs.Boss{}
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
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

			bosses = append(bosses, boss)
		})
		fmt.Println(bosses)
	})

	c.Visit("https://darksouls3.wiki.fextralife.com/Bosses")
}
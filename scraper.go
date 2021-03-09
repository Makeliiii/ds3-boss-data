package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type boss struct {
	Boss string
	Location string
	NPCSummoning string
	Weakness string
	Resistance string
	Immunity string
	Parryable string
	Optional string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("darksouls3.wiki.fextralife.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	bosses := []boss{}
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			boss := boss {
				row.ChildText("td:nth-child(1)"),
				row.ChildText("td:nth-child(2)"),
				row.ChildText("td:nth-child(3)"),
				row.ChildText("td:nth-child(4)"),
				row.ChildText("td:nth-child(5)"),
				row.ChildText("td:nth-child(6)"),
				row.ChildText("td:nth-child(7)"),
				row.ChildText("td:nth-child(8)"),
			}

			bosses = append(bosses, boss)
		})
		fmt.Println(bosses)
	})

	c.Visit("https://darksouls3.wiki.fextralife.com/Bosses")
}
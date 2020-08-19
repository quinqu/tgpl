package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

const (
	Day = time.Hour * 24

	Year = Day * 365

	Month = Year / 12
)

type Categories struct {
	Today, ThisMonth, ThisYear, Older []github.Issue
}

func (c *Categories) ToMap() map[string][]github.Issue {
	return map[string][]github.Issue{
		"today":      c.Today,
		"this-month": c.ThisMonth,
		"this-year":  c.ThisYear,
		"older":      c.Older,
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var cats Categories
	for _, item := range result.Items {
		age := time.Since(item.CreatedAt)

		switch {
		case age < Day:
			cats.Today = append(cats.Today, item)
		case age < Month:
			cats.ThisMonth = append(cats.ThisMonth, item)
		case age < Year:
			cats.ThisYear = append(cats.ThisYear, item)
		case age > Year:
			cats.Older = append(cats.Older, item)
		}
	}

	for k, v := range cats.ToMap() {

		fmt.Printf("\n------ %s ------\n", k)
		for _, issue := range v {
			fmt.Printf("  Issue %d: %s - %s\n", issue.Number, issue.Title, issue.CreatedAt)
		}
	}
}

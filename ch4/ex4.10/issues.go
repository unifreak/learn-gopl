// Modify issues to report the results in age categories, say less than a month
// old, less than a year old, and more than a year old.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"learn/gopl/ch4/github"
)

type category []*github.Issue

func display(c category, title string) {
	if len(c) > 0 {
		fmt.Printf("\n%s:\n", title)
		for _, item := range c {
			fmt.Printf("#%-5d %v %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
		}
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	monthAgo := time.Now().AddDate(0, -1, 0)
	yearAgo := time.Now().AddDate(-1, 0, 0)

	var inMonth, inYear, beforeYear category
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(monthAgo):
			inMonth = append(inMonth, item)
		case item.CreatedAt.After(yearAgo):
			inYear = append(inYear, item)
		case item.CreatedAt.Before(yearAgo):
			beforeYear = append(beforeYear, item)
		}
	}


	display(inMonth, "Created Less Than a Month")
	display(inYear, "Created More Than a Month Ago")
	display(beforeYear, "Created More Than a Year Ago")
}


// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 144.
// Измените программу issues так, чтобы она сообщала о результатах с учетом их давности,
// деля на категории, например, поданные менее месяца назад,
// менее года назад и более года.
package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
	"sort"
	"time"
)

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.Before(result.Items[j].CreatedAt)
	})
	today := time.Now()
	month := today.Add(-30 * 24 * time.Hour)
	year := today.Add(-365 * 24 * time.Hour)
	var old []*github.Issue
	var lastMonth []*github.Issue
	var lastYear []*github.Issue
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			lastMonth = append(lastMonth, item)
		} else if item.CreatedAt.After(year) {
			lastYear = append(lastYear, item)
		} else {
			old = append(old, item)
		}
	}
	if len(old) > 0 {
		fmt.Printf("%d old issues:\n", len(old))
	}
	for _, item := range old {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}
	if len(lastYear) > 0 {
		fmt.Printf("%d issues in last year:\n", len(lastYear))
	}
	for _, item := range lastYear {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}
	if len(lastMonth) > 0 {
		fmt.Printf("%d issues in last month:\n", len(lastMonth))
	}
	for _, item := range lastMonth {
		fmt.Printf("#%-5d %9.9s %v %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */

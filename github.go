package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	seasonsCat := make(map[string][]*Issue)
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Themes: %d\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.After(time.Now().AddDate(0, 0, -30)) {
			seasonsCat["month"] = append(seasonsCat["month"], item)
		}
		if item.CreatedAt.Year() == time.Now().Year() {
			seasonsCat["year"] = append(seasonsCat["year"], item)
		}
		if item.CreatedAt.After(time.Now().AddDate(0, -6, 0)) {
			seasonsCat["half-year"] = append(seasonsCat["half-year"], item)
		}
	}
	for cat, items := range seasonsCat {
		for _, item := range items {
			fmt.Printf("#%-5d %-10.10s %-9.9s %.55s\n", item.Number, cat, item.User.Login, item.Title)
		}
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

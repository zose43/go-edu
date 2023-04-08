package main

import (
	"encoding/json"
	"fmt"
	html "html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
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
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	renderHTML(result)
	templ := reportTemplate()
	report := template.Must(
		template.New("github_issues").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ),
	)
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func reportTemplate() string {
	return `{{.TotalCount}} themes
{{range .Items}}--------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func renderHTML(res *IssuesSearchResult) {
	f, err := os.Create("github_issue.html")
	defer f.Close()
	doc := html.Must(
		html.New("github_issues_html").Funcs(html.FuncMap{"daysAgo": daysAgo}).Parse(`
<h1>{{.TotalCount}}</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
<th>Age</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
<td>{{.State}}</td>
<td>{{.User.Login}}</td>
<td>{{.Title}}</td>
<td>{{.CreatedAt | daysAgo}} days</td>
</tr>
{{end}}
</table>`),
	)
	if err = doc.Execute(f, res); err != nil {
		fmt.Printf("HTML document not created, %s\n", err)
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

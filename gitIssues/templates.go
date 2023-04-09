package gitIssues

import (
	html "html/template"
	"io"
	"log"
	"os"
	"regexp"
	"text/template"
)

type Render struct {
	All *IssuesSearchResult
	Issue *Issue
	Templ string
}

func RenderHTML(r *Render, w io.Writer) {
	re := regexp.MustCompile(`.*/`)
	name := string(re.ReplaceAll([]byte(r.Templ),[]byte("")))
	doc := html.Must(
		html.New(name).ParseFiles(r.Templ),
	)
	if r.Issue != nil {
		if err := doc.Execute(w, r.Issue); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := doc.Execute(w, r.All); err != nil {
			log.Fatal(err)
		}
	}
}

func RenderTemplate(res *IssuesSearchResult) {
	templ := `{{.TotalCount}} themes
{{range .Items}}--------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.DaysAgo}} days
{{end}}`
	report := template.Must(
		template.New("github_issues").Parse(templ),
	)
	if err := report.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}

package sorting

import (
	"fmt"
	html "html/template"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Album  string
	Title  string
	Artist string
	Year   int
	Length time.Duration
}

type CustomSort struct {
	T     []*Track
	Cless func(x, y *Track) bool
	Type  string
}

func (cs *CustomSort) Less(i, j int) bool {
	return cs.Cless((*cs).T[i], (*cs).T[j])
}

func (cs *CustomSort) Len() int {
	return len((*cs).T)
}

func (cs *CustomSort) Swap(i, j int) {
	(*cs).T[i], (*cs).T[j] = (*cs).T[j], (*cs).T[i]
}

var Tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(t string) time.Duration {
	d, err := time.ParseDuration(t)
	if err != nil {
		panic(t)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 5, 4, ' ', 0)
	fmt.Fprintf(tw, format, "Album", "Title", "Artist", "Year", "Length")
	fmt.Fprintf(tw, format, "------", "------", "------", "------", "------")
	for _, track := range tracks {
		fmt.Fprintf(tw, format, track.Album, track.Title, track.Artist, track.Year, track.Length)
	}
	if err := tw.Flush(); err != nil {
		panic(tw)
	}
}

func PrintTracksHTML(w http.ResponseWriter, r *http.Request) error {
	const msg = "Can't print HTML %w\n"
	tr := CustomSort{T: Tracks, Cless: nil, Type: "asc"}

	q := r.URL.Query()
	switch q.Get("s") {
	case "asc":
		tr.Cless = compare()
		tr.Type = "desc"
		sort.Sort(&tr)
	case "desc":
		tr.Cless = compare()
		tr.Type = "asc"
		sort.Sort(sort.Reverse(&tr))
	}

	t, err := html.New("index.html").ParseFiles("sorting/view/index.html")
	if err != nil {
		return fmt.Errorf(msg, err)
	}
	if err := t.Execute(w, tr); err != nil {
		return fmt.Errorf(msg, err)
	}
	return nil
}

func compare() func(x, y *Track) bool {
	return func(x, y *Track) bool {
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return false
	}
}

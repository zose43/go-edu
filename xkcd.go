package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Resp struct {
	Month string
	Id    int `json:"num"`
	Year  string
	Title string
	Img   string
}

var data = map[string]Resp{}

func main() {
	fmt.Println("Type a title in 1 row")
	t, err := bufio.NewReader(os.Stdin).ReadString('\n')
	t = strings.TrimSpace(t)
	if err != nil {
		log.Fatalf("Invalid comics title %s", err)
	}
	if story, ok := data[t]; !ok {
		fmt.Printf("%s not in lists", t)
	} else {
		fmt.Printf("#%d %s %s %s", story.Id, story.Year, story.Title, story.Month)
	}
}

func init() {
	i := 1
	for {
		url := "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
		b := new(Resp)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Invalid resp:%v\n", err)
			resp.Body.Close()
			break
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()
			break
		}
		if err := json.NewDecoder(resp.Body).Decode(b); err != nil {
			fmt.Printf("Invalid decode %s\n", err)
			continue
		}
		data[b.Title] = *b
		i++
	}
}

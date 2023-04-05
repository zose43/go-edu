package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const URL = "https://api.github.com/repos/zose43/go-edu/issues"

type GHIssue struct {
	Owner  string   `json:"owner"`
	Repo   string   `json:"repo"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
	State  string   `json:"state"`
	Sort   string   `json:"sort"`
}

type GHResponse struct {
	Url    string `json:"html_url"`
	Id     int
	Number int
	Title  string
	Body   string
	State  string
}

func main() {
	readENV()
	id, act := mustAction()
	switch *act {
	case "create":
		createGH()
	case "delete":
		deleteGH(id)
	case "update":
		updateGH(id)
	case "list":
		listGH()
	default:
		fmt.Printf("Undefined action given\n")
	}
}

func mustAction() (*uint, *string) {
	action := flag.String("a", "", "Action for github-issues-crud")
	id := flag.Uint("id", 0, "Issue id for update/delete actions")
	flag.Parse()
	if *action == "" && isValidAction(action) {
		log.Fatal("Invalid action value")
	}
	if *action == "delete" && *id == 0 || *action == "update" && *id == 0 {
		log.Fatal("Id must be greater then 0!")
	}
	return id, action
}

func isValidAction(a *string) bool {
	usages := [...]string{"create", "delete", "update", "list"}
	for _, v := range usages {
		if v == *a {
			return true
		}
	}
	return false
}

func request(method string, body io.Reader, url string) *http.Request {
	req, err := http.NewRequest(method, URL+url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GITHUB_ISSUES_TOKEN"))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	return req
}

func createGH() {
	body := new(GHIssue)
	fill(body)
	j, err := json.Marshal(&body)
	if err != nil {
		log.Fatal(err)
	}
	req := request("POST", bytes.NewReader(j), "")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		fmt.Printf("Issue succesfully created")
	}
}

func fill(issue *GHIssue) {
	issue.Owner = "zose43"
	issue.Repo = "go-edu"
	fmt.Printf("Type a Titile\n")
	issue.Title, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	issue.Title = strings.TrimSpace(issue.Title)
	fmt.Printf("Type a Labels\n")
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		issue.Labels = append(issue.Labels, scan.Text())
	}
	fmt.Printf("Type a Body\n")
	cmd := exec.Command("xed")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	issue.Body, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	issue.Body = strings.TrimSpace(issue.Body)
}

func deleteGH(id *uint) {
	body := GHIssue{
		Owner: "zose43",
		Repo:  "go-edu",
	}
	j, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	req := request("DELETE", bytes.NewReader(j), "/"+strconv.Itoa(int(*id))+"/lock")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		fmt.Printf("Succesfully deleted #%d", int(*id))
	}
}

func listGH() {
	body := GHIssue{
		Owner: "zose43",
		Repo:  "go-edu",
		State: "open",
	}
	fmt.Printf("Type a state(open/closed)\n")
	body.State, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("Type a sort by state\n")
	body.Sort, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	j, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	req := request("GET", bytes.NewReader(j), "")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		data := make([]GHResponse, 0)
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			log.Fatal(err)
		}
		for _, item := range data {
			fmt.Printf("id: %d #%d %s %s\n", item.Id, item.Number, item.Title, item.Body)
		}
	} else {
		fmt.Printf("Code=%d, Status=%s\n", resp.StatusCode, resp.Status)
	}
}

func updateGH(id *uint) {
	body := new(GHIssue)
	fill(body)
	j, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	req := request("PATCH", bytes.NewReader(j), "/"+strconv.Itoa(int(*id)))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Printf("Succesfully updated #%d", int(*id))
	}
}

func readENV() {
	f, err := os.Open(".env")
	if err != nil {
		log.Fatal("Cannot read ENV file")
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		v := scan.Text()
		if err := os.Setenv(v[:strings.Index(v, "=")], v[strings.Index(v, "=")+1:]); err != nil {
			log.Fatal(err)
		}
	}
}

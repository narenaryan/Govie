// Go program to read an integer from STDIN and output it to STDOUT
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

const (
	imdbURL = "http://www.imdb.com/title/"
	omdbURL = "http://www.omdbapi.com/"
)

type movie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	ImdbId     string `json:"imdbID"`
	Runtime    string `json:"Runtime"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	Plot       string `json:"Plot"`
	Actors     string `json:"Actors"`
}

func MakeRequest(url string, ch chan<- string) {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%s", body)
}

func BuildRequest(key string, year string) string {
	u, _ := url.Parse(omdbURL)
	q := u.Query()
	if year != "" {
		q.Set("y", year)
	}
	q.Set("t", key)
	q.Set("r", "json")
	q.Set("plot", "short")
	u.RawQuery = q.Encode()
	return u.String()
}

func main() {
	openMovie := flag.Bool("o", false, "open movie page/s in IMDB!")
	printdetails := flag.Bool("d", false, "lists the details of the movie/s!")
	includeYear := flag.String("y", "", "Narrows result to the given year")
	flag.Parse()

	flag_len := len(flag.Args())
	flag_arr := flag.Args()
	movieList := GetMovieList(flag_len, flag_arr, *includeYear)

	switch {
	case *openMovie:
		var err error
		for _, movie := range movieList {
			switch runtime.GOOS {
			case "linux":
				err = exec.Command("xdg-open", imdbURL+movie.ImdbId).Start()
			case "windows", "darwin":
				err = exec.Command("open", imdbURL+movie.ImdbId).Start()
			default:
				err = fmt.Errorf("unsupported platform")
			}
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	case *printdetails:
		for _, movie := range movieList {
			fmt.Printf("%s\n", strings.Repeat("<", 2*len(movie.Title)))
			fmt.Println("Movie:")
			fmt.Printf("%3s(%s)\n\n", movie.Title, movie.Year)
			fmt.Println("IMDB Rating:")
			fmt.Printf("%3s\t\n\n", movie.ImdbRating)
			fmt.Println("Plot:")
			fmt.Printf("%3s\t\n\n", movie.Plot)
			fmt.Printf("%s\n", strings.Repeat(">", 2*len(movie.Title)))
		}
	}
	fmt.Print()
}

func GetMovieList(flag_len int, flag_arr []string, year string) []movie {
	movieList := []movie{}
	ch := make(chan string)
	movieLen := flag_len
	for _, key := range flag_arr {
		go MakeRequest(BuildRequest(key, year), ch)
	}

	for i := 0; i < movieLen; i++ {
		var m movie
		json.Unmarshal([]byte(<-ch), &m)
		movieList = append(movieList, m)
	}
	return movieList
}

// Go program to read an integer from STDIN and output it to STDOUT
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	"sync"
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
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	ch <- fmt.Sprintf("%s", body)
}

func fetchImage(name string, url string, path string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Use mX300 for getting good quality poster
	if url != "N/A" {
		qualityURL := strings.Replace(url, "V1_SX300.jpg", "V1_MX300.jpg", 1)
		response, err := http.Get(qualityURL)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()
		file, err := os.Create(path + "/" + name + ".jpg")
		if err != nil {
			log.Println("Unable to create file at given directory path:", path)
		}
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		log.Printf("Successfully downloaded %s poster to %s", name, path)
	} else {
		log.Println("Poster not found for movie: ", name)
	}
}

func BuildRequest(key string, year string) string {
	u, err := url.Parse(omdbURL)
	if err != nil {
		log.Println(err.Error())
	}
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
		log.Println("Successfully Fetched movie details for movie:", m.Title)
	}
	return movieList
}

func main() {
	var wg sync.WaitGroup
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("error fetching user: %v", err)
	}
	logPath := usr.HomeDir
	f, err := os.OpenFile(logPath+"/govie.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	openMovie := flag.Bool("o", false, "open movie page/s in IMDB!")
	printdetails := flag.Bool("d", false, "lists the details of the movie/s!")
	includeYear := flag.String("y", "", "Narrows result to the given year!")
	singlePoster := flag.Bool("p", false, "Downloads poster of a movie!")
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
				log.Println(err.Error())
			} else {
				log.Println("Successfully opened IMDB page for movie:", movie.Title)
			}
		}

	case *printdetails:
		for _, movie := range movieList {
			fmt.Printf("%s\n", strings.Repeat(">", 3*len(movie.Title)))
			fmt.Println("Movie:")
			fmt.Printf("\t%3s(%s)\n\n", movie.Title, movie.Year)
			fmt.Println("IMDB Rating:")
			fmt.Printf("\t%3s\t\n\n", movie.ImdbRating)
			fmt.Println("MeteCritic Score:")
			fmt.Printf("\t%s\t\n\n", movie.Metascore)
			fmt.Println("Plot:")
			fmt.Printf("\t%3s\t\n\n", movie.Plot)
			fmt.Printf("%s\n", strings.Repeat("<", 3*len(movie.Title)))
		}

	case *singlePoster:
		if flag_len < 2 {
			fmt.Println("Please mention the download directory! (ex: ./govie -p movie download_dir)")
		} else {
			_path := flag_arr[flag_len-1]
			_movies := flag_arr[:flag_len-1]
			wg.Add(flag_len - 1)
			log.Println("Started downloading poster to: ", _path)
			movieList = GetMovieList(flag_len-1, _movies, "")
			for _, movie := range movieList {
				go fetchImage(movie.Title, movie.Poster, _path, &wg)
			}
			fmt.Print("")
		}
	}
	wg.Wait()
	fmt.Print("")
}

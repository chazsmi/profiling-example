package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	port := flag.String("port", "8080", "Server port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		food := selectFoods()
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(200)
		w.Write([]byte(food))
	})

	// Starts the web server
	http.ListenAndServe(":"+*port, nil)
}

func selectFoods() string {
	food := []string{
		"Salad",
		"chips",
		"ketchup",
		"icecream",
		"apples",
		"bananas",
		"curry",
		"milk",
		"eggs",
	}
	feedme := make(chan string)
	for i := 0; i < 100; i++ {
		go func() {
			feedme <- food[random(1, len(food))]
		}()
	}
	c := 0
	yourfood := []string{}
	for {
		yourfood = append(yourfood, <-feedme)
		c++
		if len(yourfood) > 2 {
			return fmt.Sprintf("You have %s\n", yourfood)
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

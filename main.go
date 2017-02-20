package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile).Stop()
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
			fmt.Println("You have %s", yourfood)
			break
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

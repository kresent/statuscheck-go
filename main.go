package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

func main() {
	links := []string{
		"http://sales.homeapp.ru/api/v2/callOperator/phone",
		"http://homeapp.ru/search_new/map",
		"http://crm.homeapp.ru/",
		"http://test.homeapp.ru/",
		"http://homeapp.ru/",
		"http://nonworkinglink.ru/",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		color := color.New(color.FgRed)
		color.Println(link, "might be down")

		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}

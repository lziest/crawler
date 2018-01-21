package main

import (
	"flag"
	"fmt"
	"github.com/ericchiang/css"
	"github.com/robfig/cron"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

var (
	devMode bool
)

func main() {
	flag.BoolVar(&devMode, "dev", false, "dev mode")
	flag.Parse()
	c := cron.New()
	c.Start()
	if devMode {
		c.AddFunc("0-59/5 0-59 * * * *", RunQueries)
	} else {
		c.AddFunc("0 0-59 * * * *", RunQueries)
	}
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func RunQueries() {
	QueryHenryHub()
}

type NGIDailyChange struct {
	Price      string
	Percentage string
}

func QueryHenryHub() {
	resp, err := http.Get("http://www.naturalgasintel.com/data/data_products/daily?region_id=south-louisiana&location_id=SLAHH")
	if err != nil {
		log.Printf("failed to fetch Henry Hub: %v", err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("failed to parse Henry Hub resp: %v", err)
	}
	sel, err := css.Compile("td.numeric > div > span")
	if err != nil {
		panic(err)
	}
	NewDailyChange := NGIDailyChange{}
	for i, ele := range sel.Select(doc) {
		// only need the first two values
		if i == 0 {
			NewDailyChange.Price = ele.FirstChild.Data
		} else if i == 1 {
			NewDailyChange.Percentage = ele.FirstChild.Data
		}
	}
	fmt.Println(NewDailyChange)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

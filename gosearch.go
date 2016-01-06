package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"google.golang.org/api/customsearch/v1"
)

const ApiKey = "AIzaSyAgMDxNmzb82S7k6-tNi6vj5GOiOjR6HlE"
const Cx = "005430158042492320947:hww-v4cc4ow"

func main() {
	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=", ApiKey, Cx)
	fmt.Println("Enter a search term:")
	var q string
	fmt.Scanln(&q)
	resp, err := http.Get(url + q)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var r customsearch.Search
	if err := json.Unmarshal(body, &r); err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(r.Items); i++ {
		res := r.Items[i]
		fmt.Printf("[%v] %s (%s)\n", i+1, res.Title, res.Link)
	}
}

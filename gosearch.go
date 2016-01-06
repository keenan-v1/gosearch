package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/customsearch/v1"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const ApiKey = "AIzaSyAgMDxNmzb82S7k6-tNi6vj5GOiOjR6HlE"
const Cx = "005430158042492320947:hww-v4cc4ow"

func main() {
	sUrl := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=", ApiKey, Cx)
	var q string
	if len(os.Args) > 1 {
		q = strings.Join(os.Args[1:], " ")
		fmt.Println("Searching for:", q)
	} else {
		fmt.Println("Enter a search term:")
		fmt.Scanln(&q)
	}
	resp, err := http.Get(sUrl + url.QueryEscape(q))
	if err != nil {
		log.Println("Http:Get", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll", err)
		return
	}
	var r customsearch.Search
	if err := json.Unmarshal(body, &r); err != nil {
		log.Println("json:Unmarshal", err)
		return
	}
	for i := 0; i < len(r.Items); i++ {
		res := r.Items[i]
		fmt.Printf("[%v] %s (%s)\n", i+1, res.Title, res.Link)
	}
}

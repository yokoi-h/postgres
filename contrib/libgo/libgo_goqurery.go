package main

import (
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"log"
)

//export jpy
func jpy() float32 {
	var cur float32 = 0.0

	url := "https://www.google.com/finance/converter?a=1&from=USD&to=JPY"
	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Println(err)
		return cur
	}

	doc.Find("div#currency_converter_result").Each(func(_ int, s *goquery.Selection) {
		line := s.Text()
		l := strings.Split(line, " ")
		v, _ := strconv.ParseFloat(l[3], 32)
		cur = float32(v)
		log.Println(v)
	})

	return float32(cur)
}

func init() {
	log.Println("loaded.")
}

func main() {
	c := jpy()
	log.Println(c)
}

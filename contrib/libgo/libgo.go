package main

import (
	"C"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

//export jpy
func jpy() float32 {
	var cur float32 = 0.0
	resp, err := http.Get("https://www.google.com/finance/converter?a=1&from=USD&to=JPY")
	if err != nil {
		fmt.Println(err)
		return cur
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	var f func(*html.Node)
	f = func(n *html.Node) {
		fmt.Println(n.Data)
		i := strings.Index(n.Data, "JPY")
		if i != -1 {
			a := strings.Split(n.Data, " ")
			v, _ := strconv.ParseFloat(a[0], 32)
			cur = float32(v)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if cur == 0.0 {
				f(c)
			}
		}
	}

	f(doc)
	return float32(cur)
}

func init() {
	log.Println("loaded.")
}

func main() {
	c := jpy()
	fmt.Println(c)
}

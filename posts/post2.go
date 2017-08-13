package posts

import (
	"log"
	"net/http"
)

func loadPost2() {
	// `d`, the reference to the DOM is declared in a previous post
	d.GetElementByID("post-2-title").SetInnerHTML("Day 2: Setting up a REST API")
	d.GetElementByID("post-2-attribution").SetInnerHTML("By Joseph Choi under <a class='post-category post-category-js'>Go</a> <a class='post-category post-category-js'>REST API</a> <a class='post-category post-category-js'>Heroku</a>")
	d.GetElementByID("post-2-description-1").SetInnerHTML("In an attempt to get a fullstack Golang project up and running, I decided to deploy a very simple Golang REST API to Heroku, and connect it to the GopherJS frontend. Below is the relevant server-side and frontend code.")
	resp, err := http.Get("https://gopherjs-api.herokuapp.com/weatherreport")
	if err != nil {
		log.Fatal(err)
	}
	println("%s\n", resp)
}

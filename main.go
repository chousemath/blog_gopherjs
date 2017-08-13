package main

import "github.com/gopherjs/gopherjs/js"
import "honnef.co/go/js/dom"

var canvasPost1Img1 = "img/common/gopher-1.png"
var canvasPost1Img2 = "img/common/gopher-2.png"

func main() {
	d := dom.GetWindow().Document()

	d.GetElementByID("blog-title").SetInnerHTML("Joseph's Golang Blog")
	d.GetElementByID("blog-subtitle").SetInnerHTML("Learning the Ways of the Gopher")

	d.GetElementByID("post-1-title").SetInnerHTML("Day 1: Setting up GopherJS")
	d.GetElementByID("post-1-attribution").SetInnerHTML("By Joseph Choi under <a class='post-category post-category-js'>Go</a>")
	d.GetElementByID("post-1-description-1").SetInnerHTML("Today was my first day with GopherJS. In order to drive myself to learn this new language, I am hoping to make short, Go-related posts on a regular basis. Wish me luck!")
	// fetch the canvas element
	canvasPost1 := d.GetElementByID("post-1-canvas-1")
	canvasPost1.SetAttribute("style", "border: 3px solid #000000;")
	// extract the 2D context
	contextPost1 := canvasPost1.Underlying().Call("getContext", "2d")
	// initialize a new image
	post1Img1 := js.Global.Get("Image").New()
	// define the src of the new image
	post1Img1.Set("src", canvasPost1Img1)
	// when the canvas is loaded, insert the image
	post1Img1.Call("addEventListener", "load", func() {
		println("Drawing image with GopherJS!")
		contextPost1.Call("drawImage", post1Img1, 0, 0)
	})

	canvasPost2 := d.GetElementByID("post-1-canvas-2")
	canvasPost2.SetAttribute("style", "border: 3px solid #0000ff;")
	contextPost2 := canvasPost2.Underlying().Call("getContext", "2d")
	post1Img2 := js.Global.Get("Image").New()
	post1Img2.Set("src", canvasPost1Img2)
	post1Img2.Call("addEventListener", "load", func() {
		println("Drawing another image with GopherJS!")
		contextPost2.Call("drawImage", post1Img2, 0, 0)
	})
}

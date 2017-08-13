package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"honnef.co/go/js/dom"
)

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// WeatherReport represents a typical weather report
type WeatherReport struct {
	ID              int32   `json:"id,omitempty"`
	Day             string  `json:"day,omitempty"`
	TemperatureLow  float32 `json:"temperatureLow,omitempty"`
	TemperatureHigh float32 `json:"temperatureHigh,omitempty"`
	Precipitation   float32 `json:"precipitation,omitempty"`
	Humidity        float32 `json:"humidity,omitempty"`
	Wind            float32 `json:"wind,omitempty"`
}

// WeatherReports represents a typical weather report
type WeatherReports struct {
	Reports []WeatherReport `json:"reports,omitempty"`
}

func formatFloat(f float32) string {
	return fmt.Sprintf("%.2f", f)
}

func appendColHeader(row dom.Element, colContent string) {
	headerCol := d.CreateElement("TH")
	headerCol.SetInnerHTML(colContent)
	row.AppendChild(headerCol)
}

func appendCol(row dom.Element, colContent string) {
	headerCol := d.CreateElement("TD")
	headerCol.SetInnerHTML(colContent)
	row.AppendChild(headerCol)
}

func loadPost2() {
	// `d`, the reference to the DOM is declared in a previous post
	d.GetElementByID("post-2-title").SetInnerHTML("Day 2: Connecting a Golang REST API")
	d.GetElementByID("post-2-attribution").SetInnerHTML("By Joseph Choi under <a class='post-category post-category-js'>Go</a> <a class='post-category post-category-js'>REST API</a> <a class='post-category post-category-js'>Heroku</a>")
	d.GetElementByID("post-2-description-1").SetInnerHTML("In an attempt to get a fullstack Golang project up and running, I decided to deploy a very simple Golang REST API to Heroku, and connect it to the GopherJS frontend. Below is the relevant server-side and frontend code.")

	weatherReportsHeader := d.GetElementByID("weather-reports-header")
	headerLabels := []string{"Day", "TempLo", "TempHi", "Precip", "Humid", "Wind"}
	headerRow := d.CreateElement("TR")
	for _, label := range headerLabels {
		appendColHeader(headerRow, label)
	}

	weatherReportsHeader.AppendChild(headerRow)

	url := "https://gopherjs-api.herokuapp.com/weatherreports"
	weatherReports := new(WeatherReports)
	getJSON(url, weatherReports)

	weatherReportsTable := d.GetElementByID("weather-reports")
	for _, weatherReport := range weatherReports.Reports {
		// create a new table row node
		row := d.CreateElement("TR")
		rowValues := []string{
			weatherReport.Day,
			formatFloat(weatherReport.TemperatureLow),
			formatFloat(weatherReport.TemperatureHigh),
			formatFloat(weatherReport.Precipitation),
			formatFloat(weatherReport.Humidity),
			formatFloat(weatherReport.Wind),
		}

		for _, value := range rowValues {
			appendCol(row, value)
		}

		weatherReportsTable.AppendChild(row)
	}
}

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/kellydunn/golang-geo"
	"github.com/go-martini/martini"
)

type geocodeResult struct {
	Address string
	Point   *geo.Point
}

func main() {


	flag.Usage = usage
	// Adding new option web for web interface
	webPtr :=flag.Bool("web", false, "web interface of the app")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	query := strings.Join(args, " ")

	result, err := geocode(query, new(geo.GoogleGeocoder))
	if err != nil {
		printErr(err)
	}

	tmpl(os.Stdout, resultTemplate, result)
}

const usageTemplate = `Geodude is a tiny command-line utility for geocoding addresses.

Usage:

  geodude [address]
  geodude -web 

`

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, nil)
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func printErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(2)
}

const resultTemplate = `Address: {{.Address}}
Coordinates: {{.Point.Lat}}, {{.Point.Lng}}

`

func geocode(query string, geocoder geo.Geocoder) (*geocodeResult, error) {
	point, err := geocoder.Geocode(query)
	if err != nil {
		return nil, err
	}

	addr, err := geocoder.ReverseGeocode(point)
	if err != nil {
		return nil, err
	}

	result := &geocodeResult{
		Address: addr,
		Point:   point,
	}

	return result, nil
}

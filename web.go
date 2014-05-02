package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "<h1>Geo Dude App</h1>"
  })
  m.Run()
}

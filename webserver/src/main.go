package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width  float64
	Length float64
}

func (r Rect) Area() float64 {
	return r.Length * r.Width
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func main() {
	var r Shape
	r = Rect{5.5, 4.5}

	a := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, fmt.Sprint(r.Area()))
	}

	p := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, fmt.Sprint(r.Perimeter()))
	}

	http.HandleFunc("/area", a)
	http.HandleFunc("/perimeter", p)
	log.Fatal(http.ListenAndServe(":80", nil))
}

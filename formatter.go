package main

import "fmt"

type Color string

var (
	Green   = Color("\033[32m")
	Magenta = Color("\033[35m")
	Red     = Color("\033[91m")
	White   = Color("\033[37m")
	Reset   = Color("\033[0m")
)

type Formatter interface {
	fmt.Stringer
	Colorer
}

type Colorer interface {
	Color() string
}

package main

type Menu struct {
	Day string
	section []Section
}

type Section struct {
	Name string
	food []Food
}
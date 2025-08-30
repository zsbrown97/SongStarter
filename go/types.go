package main

type Key struct {
	Root string
	Mode string
	Notes []Note
}

type Note struct {
	Note string
	Degree string
}

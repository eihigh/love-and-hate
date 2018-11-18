package main

type generator interface {
	update(syms []*symbol)
}

var (
	generators = []generator{}
)

func generate() {
}

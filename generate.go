package main

type generator interface {
	update(syms []*symbol)
}

var (
	generators = []generator{}
)

func generate() {

	for _, g := range generators {
		g.update(symbols)
	}
}

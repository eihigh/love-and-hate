package main

var (
	symbols []*symbol
)

type symbol struct {
	isLove bool
}

func newLove() *symbol {
	return &symbol{isLove: true}
}

func newHate() *symbol {
	return &symbol{isLove: false}
}

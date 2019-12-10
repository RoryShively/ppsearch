package main

type Engine interface {
	AddWord() error
	AddPage() error
	Clear() error
}

// type InMemoryEngine struct {

// }

// func (v *InMemoryEngine) Clear {

// }

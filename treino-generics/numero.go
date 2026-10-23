package main

type Numero interface {
	~int | ~int64 | ~float64
}

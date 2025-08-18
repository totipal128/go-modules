package main

import (
	"gitlab.com/package7225033/go-modules/check"
)

type People struct {
	Name string
	Age  string
}

func main() {
	var (
		as  People
		asd []People
	)

	check.CheckType(as)
	check.CheckType(asd)
}

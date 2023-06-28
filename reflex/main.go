package main

import (
	"reflex/scrub/pkg"

	"github.com/kr/pretty"
)

type TestData struct {
	Name     string
	Password string
	Nested   *TestData
}

func main() {
	t := &TestData{
		Name:     "test",
		Password: "abc",
		Nested: &TestData{
			Name:     "nested",
			Password: "defxyz",
			Nested: &TestData{
				Name:     "nested2",
				Password: "Привет",
			},
		},
	}

	pretty.Print(pkg.Scrub(t, map[string]string{
		"Password": "*",
	}))
}

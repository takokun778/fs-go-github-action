package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Test struct {
	Repo string `csv:"repo"`
	Name string `csv:"name"`
}

func main() {
	env1 := os.Getenv("ENV1")
	env2 := os.Getenv("ENV2")
	log.Printf("env1: %s\n", env1)
	log.Printf("env2: %s\n", env2)

	os.Rename("test.csv", "src.csv")

	// csvを読み込む
	src, err := os.OpenFile("src.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	tests := []*Test{}

	if err := gocsv.UnmarshalFile(src, &tests); err != nil { // Load clients from file
		panic(err)
	}
	for _, test := range tests {
		fmt.Printf("%s %s\n", test.Repo, test.Name)
	}

	tests = append(tests, &Test{Repo: "test", Name: "test"})

	// 差分を計算

	dst, err := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	// csvに書き込む
	if err := gocsv.MarshalFile(&tests, dst); err != nil {
		log.Fatal(err)
	}

	os.Remove("src.csv")
}

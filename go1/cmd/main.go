package main

import (
	"flag"

	worker "github.com/demo-apps/go-gin-app/go1/workers"
)

	var phrase = flag.String("word", "", "chain of word or one word")
	var urls = flag.String("sites", "", "chain of sites or one site adress")
	var file = flag.String("file", "", "path to the file")

func main() {
	flag.Parse()
	worker.Worker(*phrase, *urls, *file)
}


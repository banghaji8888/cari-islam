package main

import (
	"cari-islam/config"
	"cari-islam/crawler/rumaysho"
	"cari-islam/db"
	"cari-islam/elasticsearch"
	"cari-islam/util"
	"os"
)

func main() {
	defer util.Recover()

	// METHOD=elasticsearch go run main.go

	config.Init()
	db.Init()

	method := os.Getenv("METHOD")

	if method == "elasticsearch" {
		elasticsearch.InsertToElasticSearch()
	} else {
		rumaysho.Init(method)
	}
}

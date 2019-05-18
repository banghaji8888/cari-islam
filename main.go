package main

import (
	"cari-islam/config"
	"cari-islam/crawler/rumaysho"
	"cari-islam/db"
	"cari-islam/util"
	"os"
)

func main() {
	defer util.Recover()

	// METHOD=initKonten go run main.go

	config.Init()
	db.Init()

	method := os.Getenv("METHOD")

	rumaysho.Init(method)
}

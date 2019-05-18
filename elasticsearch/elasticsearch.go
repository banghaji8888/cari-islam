package elasticsearch

import (
	"cari-islam/db"
	"cari-islam/model"
	"cari-islam/util"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// InsertToElasticSearch - insert to elasticsearch
func InsertToElasticSearch() {
	defer util.Recover()

	db := db.GetConnection()

	kontens := []model.Konten{}
	db.Where(&model.Konten{IsDone: 0}).Order("id").Limit(1).Find(&kontens)

	var raw string
	for _, konten := range kontens {
		raw = util.CurlGet(konten.URL)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
		if err != nil {
			log.Println(err)
		}

		longDescription := doc.Find(".td-post-content").Text()
		re := regexp.MustCompile(`<img.+?src=[\"'](.+?)[\"'].+?>`)
		longDescription = re.ReplaceAllString(longDescription, "")
		re = regexp.MustCompile(`\r?\n`)

		log.Println(re.ReplaceAllString(longDescription, ""))
	}
}

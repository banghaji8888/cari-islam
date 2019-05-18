package rumaysho

import (
	"cari-islam/db"
	"cari-islam/model"
	"cari-islam/repo"
	"cari-islam/util"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var sumberID = 1

//Init - filter method
func Init(method string) {
	if method == "initKonten" {
		initKonten()
	} else if method == "updateKonten" {
		updateKonten()
	}
}

// initKonten - fill first konten
func initKonten() {
	defer util.Recover()

	db := db.GetConnection()

	kategori := []model.Kategori{}
	db.Where(&model.Kategori{SumberID: sumberID}).Find(&kategori)

	var raw string
	repo := repo.InitRepo(db)

	log.Println("Isi pertama kali tabel konten rumaysho...")
	for _, kat := range kategori {
		raw = util.CurlGet(kat.URL)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
		if err != nil {
			log.Println(err)
		}

		stringLastPage := doc.Find(".page-nav .last").Text()
		lastPage, _ := strconv.Atoi(stringLastPage)

		// iterate all pages
		for i := lastPage; i > 0; i-- {
			url := strings.Join([]string{
				kat.URL,
				"page",
				strconv.Itoa(i),
			}, "/")

			log.Println(url)

			raw = util.CurlGet(url)

			docs, errs := goquery.NewDocumentFromReader(strings.NewReader(raw))
			if errs != nil {
				log.Println(errs)
			}

			docs.Find(".td-ss-main-content .td_module_10").Each(func(j int, s *goquery.Selection) {
				title := s.Find(".item-details h3").Text()
				link, _ := s.Find(".item-details h3 a").Attr("href")
				tanggal := s.Find(".td-post-date").Text()
				shortDesc := strings.TrimSpace(s.Find(".td-excerpt").Text())

				repo.SaveToKonten("rumaysho.com", title, link, tanggal, shortDesc, kat.ID)
			})
		}
	}
}

// updateKonten - update konten
func updateKonten() {
	defer util.Recover()

	db := db.GetConnection()

	var lastKonten model.Konten
	kategori := []model.Kategori{}
	db.Where(&model.Kategori{SumberID: sumberID}).Find(&kategori)

	var raw string
	repo := repo.InitRepo(db)

	log.Println("Update tabel konten rumaysho...")

	for _, kat := range kategori {
		lastKonten = model.Konten{}
		db.Where(&model.Konten{KategoriID: kat.ID}).Last(&lastKonten)
		raw = util.CurlGet(kat.URL)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
		if err != nil {
			log.Println(err)
		} else {
			doc.Find(".td-ss-main-content .td_module_10").Each(func(j int, s *goquery.Selection) {
				title := s.Find(".item-details h3").Text()
				link, _ := s.Find(".item-details h3 a").Attr("href")
				tanggal := s.Find(".td-post-date").Text()
				shortDesc := strings.TrimSpace(s.Find(".td-excerpt").Text())

				if link != lastKonten.URL {
					repo.SaveToKonten("rumaysho.com", title, link, tanggal, shortDesc, kat.ID)
				}
			})
		}
	}
}

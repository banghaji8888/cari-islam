package repo

import (
	"cari-islam/model"
	"cari-islam/util"

	"github.com/jinzhu/gorm"
)

//repo - Repo struct
type repo struct {
	conn *gorm.DB
}

//InitRepo - init repo
func InitRepo(dbconn *gorm.DB) *repo {
	return &repo{
		conn: dbconn,
	}
}

func (r *repo) SaveToKonten(title, url, tanggal, shortDesc string, kategoriID int) {
	defer util.Recover()

	konten := &model.Konten{
		Title:      title,
		URL:        url,
		Tanggal:    tanggal,
		ShortDesc:  shortDesc,
		KategoriID: kategoriID,
	}

	r.conn.Create(konten)
}

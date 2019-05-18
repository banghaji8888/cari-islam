package model

import "time"

//Konten - model
type Konten struct {
	ID          int        `gorm:"primary_key" json:"id"`
	Title       string     `json:"title" gorm:"size:200"`
	URL         string     `json:"url" gorm:"size:80000"`
	ShortDesc   string     `json:"short_desc" gorm:"size:80000"`
	Tanggal     string     `json:"tanggal" gorm:"size:20"`
	KategoriID  int        `json:"kategori_id"`
	IsDone      int        `json:"is_done"`
	CreatedDate *time.Time `json:"created_date"`
}

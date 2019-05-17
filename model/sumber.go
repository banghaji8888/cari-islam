package model

//Sumber - model
type Sumber struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Nama string `json:"nama" gorm:"size:200"`
	URL  string `json:"url" gorm:"size:80000"`
}

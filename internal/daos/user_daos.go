package daos

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id            int64  `gorm:"primarykey"`
	Nama          string `gorm:"type:varchar(255)"`
	Kata_Sandi    string `gorm:"type:varchar(255)"`
	Notelp        string `gorm:"type:varchar(255);unique"`
	Tanggal_lahir string `gorm:"type:date"`
	Jenis_kelamin string `gorm:"type:varchar(255)"`
	Tentang       string `gorm:"type:text"`
	Pekerjaan     string `gorm:"type:varchar(255)"`
	Email         string `gorm:"type:varchar(255)"`
	Id_provinsi   string `gorm:"type:varchar(255)"`
	Id_kota       string `gorm:"type:varchar(255)"`
	IsAdmin       bool `gorm:"type:boolean"`
	Updated_at    string `gorm:"type:date"`
	Created_at    string `gorm:"type:date"`
}

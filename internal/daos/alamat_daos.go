package daos

import "gorm.io/gorm"

type Alamat struct {
	gorm.Model
	Id            int64  `gorm:"primaryKey"`
	Id_user       int64  `gorm:"foreignKey:User"`
	User          []User   `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	Judul_Alamat  string `gorm:"type:varchar(255)"`
	Nama_penerima string `gorm:"type:varchar(255)"`
	No_telp       string `gorm:"type:varchar(255)"`
	Detail_alamat string `gorm:"type:varchar(255)"`
	Update_at     string `gorm:"type:date"`
	Created_at    string `gorm:"type:date"`
}

package daos

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	Id         int64  `gorm:"primaryKey"`
	Id_user    int64  `gorm:"foreignKey:User"`
	User       []User `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	Nama_toko  string `gorm:"type:varchar(255)"`
	Url_foto   string `gorm:"type:varchar(255)"`
	Update_at  string `gorm:"type:date"`
	Created_at string `gorm:"type:date"`
}

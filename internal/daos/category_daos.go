package daos

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id            int64  `gorm:"primaryKey"`
	Nama_category string `gorm:"type:varchar(255)"`
	Update_at     string `gorm:"type:date"`
	Created_at    string `gorm:"type:date"`
}

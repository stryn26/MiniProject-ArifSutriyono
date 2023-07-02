package daos

import "gorm.io/gorm"

type Log_produk struct {
	gorm.Model
	Id             int64      `gorm:"primaryKey"`
	// FK
	Id_toko        int64      `gorm:"foreignKey:User"`
	Toko           []Toko     `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	Id_category    int64      `gorm:"foreignKey:Category"`
	Category       []Category `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	// END FK
	Nama_produk    string     `gorm:"type:varchar(255)"`
	Slug           string     `gorm:"type:varchar(255)"`
	Harga_reseller string     `gorm:"type:varchar(255)"`
	Harga_konsumen string     `gorm:"type:varchar(255)"`
	Strok          int64      `gorm:"type:int"`
	Deskripsi      string     `gorm:"type:text"`
	Update_at      string     `gorm:"type:date"`
	Created_at     string     `gorm:"type:date"`
}

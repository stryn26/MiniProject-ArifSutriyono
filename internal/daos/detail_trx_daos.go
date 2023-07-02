package daos

import "gorm.io/gorm"

type Detail_trx struct {
	gorm.Model
	Id int64 `gorm:"primaryKey"`
	// FK
	Id_toko       int64        `gorm:"foreignKey:Toko"`
	Toko          []Toko       `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	Id_trx        int64        `gorm:"foreignKey:Category"`
	Trx           []Trx        `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	Id_log_produk int64        `gorm:"foreignKey:Log_produk"`
	Log_produk    []Log_produk `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	// END FK
	Kuantitas   string `gorm:"type:int"`
	Harga_total string `gorm:"type:int"`
	Update_at   string `gorm:"type:date"`
	Created_at  string `gorm:"type:date"`
}

package daos

import "gorm.io/gorm"

type Foto_produk struct {
	gorm.Model
	Id         int64    `gorm:"primaryKey"`
	Id_produk  int64    `gorm:"foreignKey:Produk"`
	Produk     []Produk `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	Update_at  string   `gorm:"type:date"`
	Created_at string   `gorm:"type:date"`
}

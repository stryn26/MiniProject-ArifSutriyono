package daos

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	Id int64 `gorm:"primaryKey"`
	// FK
	Id_user int64  `gorm:"foreignKey:User"`
	User    []User `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
	// END FK
	Alamat_pengiriman string `gorm:"type:varchar(255)"`
	Harga_total       int    `gorm:"type:int"`
	Kode_invoice      string `gorm:"type:varchar(255)"`
	Method_bayar      string `gorm:"type:int"`
	Update_at         string `gorm:"type:date"`
	Created_at        string `gorm:"type:date"`
}

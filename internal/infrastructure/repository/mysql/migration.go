package mysql

import (
	"fmt"
	"github.com/stryn26/MiniProjectEvermosRakamin/internal/daos"
	"github.com/stryn26/MiniProjectEvermosRakamin/internal/helper"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&daos.User{},
		&daos.Toko{},
		&daos.Produk{},
		&daos.Category{},
		&daos.Foto_produk{},
		&daos.Log_produk{},
		&daos.Trx{},
		&daos.Detail_trx{},
		&daos.Alamat{},
	)

	var count int64
	if mysqlDB.Migrator().HasTable(&daos.User{}) {
		mysqlDB.Model(&daos.User{}).Count(&count)
		if count < 1 {
			mysqlDB.CreateInBatches(userSeed, len(userSeed))
		}
	}

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}
package database

import (
	discountsRepo "consignku/drivers/databases/discounts"
	productTypesRepo "consignku/drivers/databases/product_types"
	productUsedTimesRepo "consignku/drivers/databases/product_used_times"
	productsRepo "consignku/drivers/databases/products"
	indonesiaCityLocationRepo "consignku/drivers/databases/thirdparties/indonesia_city_location"
	transactionsRepo "consignku/drivers/databases/transactions"
	transactionsItemRepo "consignku/drivers/databases/transactions_item"
	usersRepo "consignku/drivers/databases/users"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal((err))
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {

	db.AutoMigrate(
		&usersRepo.Users{},
		&discountsRepo.Discounts{},
		&productTypesRepo.ProductTypes{},
		&productUsedTimesRepo.ProductUsedTimes{},
		&productsRepo.Products{},
		&transactionsRepo.Transactions{},
		&transactionsItemRepo.TransactionsItem{},
		&indonesiaCityLocationRepo.City{},
	)
}

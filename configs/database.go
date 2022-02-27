package configs

import (
	"car-booking/models"
	"car-booking/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", utils.LoadEnv("DB_USERNAME"), utils.LoadEnv("DB_PASSWORD"), utils.LoadEnv("DB_HOSTNAME"), utils.LoadEnv("DB_PORT"), utils.LoadEnv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}, &models.Role{}, &models.Menu{}, &models.UserInfo{}, &models.RefreshToken{}, &models.Agency{}, &models.BillPayment{}, &models.Booking{}, &models.BranchAddress{}, &models.DiscountType{}, &models.Discount{}, &models.PaymentType{}, &models.VehicleType{}, &models.VehicleManufacturer{}); err != nil {
		panic("Cant Migrate Database")
	}

	DB = db

	return db, nil
}

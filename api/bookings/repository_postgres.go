package bookings

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type bookingRepositoryPostgres struct {
	db *gorm.DB
}

func NewBookingRepositoryPostgres(db *gorm.DB) BookingRepository {
	err := db.AutoMigrate(&BookingQuestion{}, &BookingSkill{}, &Booking{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to postgres in booking service!")

	return &bookingRepositoryPostgres{
		db: db,
	}
}

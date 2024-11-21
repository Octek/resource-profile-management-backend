package bookings

import (
	"fmt"
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

func (repo *bookingRepositoryPostgres) AddBooking(request AddBookingRequest, userId uint) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		for _, bookingReq := range request.BookingRequest {
			bookingModel := Booking{
				UserID:          userId,
				BookingDateTime: bookingReq.BookingDateTime,
				MeetingLink:     bookingReq.MeetingLink,
			}

			if err := tx.Create(&bookingModel).Error; err != nil {
				return fmt.Errorf("failed to create booking: %w", err)
			}

			var bookingSkillArray []BookingSkill
			for _, skillID := range bookingReq.SkillID {
				bookingSkill := BookingSkill{
					SkillID:   skillID,
					BookingID: bookingModel.ID,
				}
				bookingSkillArray = append(bookingSkillArray, bookingSkill)
			}

			if len(bookingSkillArray) > 0 {
				if err := tx.Create(&bookingSkillArray).Error; err != nil {
					return fmt.Errorf("failed to create booking skills for booking ID %d: %w", bookingModel.ID, err)
				}
			}

			var bookingQuestionArray []BookingQuestion
			for _, questionOptionID := range bookingReq.QuestionOptionID {
				bookingQuestion := BookingQuestion{
					QuestionOptionID: questionOptionID,
					BookingID:        bookingModel.ID,
				}
				bookingQuestionArray = append(bookingQuestionArray, bookingQuestion)
			}

			if len(bookingQuestionArray) > 0 {
				if err := tx.Create(&bookingQuestionArray).Error; err != nil {
					return fmt.Errorf("failed to create booking questions for booking ID %d: %w", bookingModel.ID, err)
				}
			}
		}

		fmt.Println("Booking, BookingSkill, and BookingQuestion records have been created successfully.")
		return nil
	})

	return err
}

func (repo *bookingRepositoryPostgres) GetBookingById(id uint) (*Booking, error) {
	var booking Booking
	if err := repo.db.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (repo *bookingRepositoryPostgres) GetUserBookingByUserIdAndBookingId(userId, BookingId uint) (*Booking, error) {
	var booking Booking
	if err := repo.db.Where("user_id = ? AND id = ?", userId, BookingId).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (repo *bookingRepositoryPostgres) UpdateBooking(Booking *Booking) error {
	return repo.db.Save(Booking).Error
}

func (repo *bookingRepositoryPostgres) GetAllUserBookingList(bookingID, userID uint) (Booking, error) {
	var booking Booking

	err := repo.db.Model(&Booking{}).
		Where("bookings.id = ? AND bookings.user_id = ?", bookingID, userID).
		Joins("JOIN booking_skills ON booking_skills.booking_id = bookings.id").
		Where("booking_skills.booking_id = ?", bookingID).
		Joins("JOIN booking_questions ON booking_questions.booking_id = bookings.id").
		Where("booking_questions.booking_id = ?", bookingID).
		Preload("QuestionOptions").
		First(&booking).
		Error

	return booking, err
}

func (repo *bookingRepositoryPostgres) DeleteUserBookingByID(id uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("booking_id = ?", id).Delete(&BookingSkill{}).Error; err != nil {
			return err
		}
		if err := tx.Where("booking_id = ?", id).Delete(&BookingQuestion{}).Error; err != nil {
			return err
		}
		result := tx.Delete(&Booking{}, id)
		if result.RowsAffected == 0 {
			return fmt.Errorf("no Booking record found for id: %d", id)
		}

		return result.Error
	})
}

func (repo *bookingRepositoryPostgres) DeleteUserBookingByUserID(userId uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		var bookingIDs []uint
		if err := tx.Model(&Booking{}).Where("user_id = ?", userId).
			Pluck("id", &bookingIDs).Error; err != nil {
			return err
		}
		if len(bookingIDs) == 0 {
			return fmt.Errorf("no Booking records found for user_id: %d", userId)
		}
		err := tx.Where("booking_id IN (?)", bookingIDs).Delete(&BookingQuestion{}).Error
		err = tx.Where("booking_id IN (?)", bookingIDs).Delete(&BookingSkill{}).Error

		err = tx.Where("id IN (?)", bookingIDs).Delete(&Booking{}).Error

		return err
	})
}

func (repo *bookingRepositoryPostgres) GetAllUserBooking(userId uint, limit int, offset int, orderBy string) ([]Booking, uint, error) {
	var BookingIDs []uint
	var exp []Booking
	var total int64

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Booking{}).
			Where("user_id = ?", userId).
			Pluck("id", &BookingIDs).Error; err != nil {
			return err
		}
		if len(BookingIDs) == 0 {
			return fmt.Errorf("no Booking records found for user_id: %d", userId)
		}

		query := tx.Model(&Booking{}).Where("deleted_at IS NULL").Where("id IN (?)", BookingIDs).Preload("QuestionOptions")
		err := query.Count(&total).Error
		err = query.Order(orderBy).Limit(limit).Offset(offset).Find(&exp).Error

		return err
	})

	return exp, uint(total), err
}

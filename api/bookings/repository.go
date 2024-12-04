package bookings

// BookingRepository Used to store and retrieve user bookings
type BookingRepository interface {
	AddBooking(booking AddBookingRequest, userId uint) error
	GetBookingById(id uint) (*Booking, error)
	GetUserBookingByUserIdAndBookingId(userId, bookingId uint) (*Booking, error)
	UpdateBooking(booking *Booking) error
	GetAllUserBookingList(bookingID, userID uint) (Booking, error)
	DeleteUserBookingByID(id uint) error
	DeleteUserBookingByUserID(id uint) error
	GetAllUserBooking(userId uint, limit int, offset int, orderBy string) ([]Booking, uint, error)
	AddBookingSkills(skillIds []uint, bookingId uint) error
	AddBookingQuestionOption(questionOptId []uint, bookingId uint) error
	//createCategories(jsonData []Category) error
}

package bookings

type BookingService struct {
	bookingRepository BookingRepository
}

func NewService(r BookingRepository) BookingService {
	return BookingService{bookingRepository: r}
}

func (svc *BookingService) AddBooking(booking *Booking, questionOptId, skillId uint) (*Booking, error) {
	return svc.bookingRepository.AddBooking(booking, questionOptId, skillId)
}

func (svc *BookingService) GetBookingById(id uint) (*Booking, error) {
	return svc.bookingRepository.GetBookingById(id)
}

func (svc *BookingService) GetUserBookingByUserIdAndBookingId(userId, bookingId uint) (*Booking, error) {
	return svc.bookingRepository.GetUserBookingByUserIdAndBookingId(userId, bookingId)
}

func (svc *BookingService) UpdateBooking(booking *Booking) error {
	return svc.bookingRepository.UpdateBooking(booking)
}

func (svc *BookingService) GetAllUserBookingList(bookingID, userID uint) (Booking, error) {
	return svc.bookingRepository.GetAllUserBookingList(bookingID, userID)
}

func (svc *BookingService) DeleteUserBookingByID(id uint) error {
	return svc.bookingRepository.DeleteUserBookingByID(id)
}

func (svc *BookingService) DeleteUserBookingByUserID(id uint) error {
	return svc.bookingRepository.DeleteUserBookingByUserID(id)
}
func (svc *BookingService) GetAllUserBooking(userId uint, limit int, offset int, orderBy string) ([]Booking, uint, error) {
	return svc.bookingRepository.GetAllUserBooking(userId, limit, offset, orderBy)
}

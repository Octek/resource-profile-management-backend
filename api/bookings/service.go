package bookings

type BookingService struct {
	bookingRepository BookingRepository
}

func NewService(r BookingRepository) BookingService {
	return BookingService{bookingRepository: r}
}

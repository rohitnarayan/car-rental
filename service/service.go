package service

import (
	"errors"

	"github.com/rohitnarayan/car-rental/model"
	"github.com/rohitnarayan/car-rental/repository"
)

type CarService interface {
	SearchCars(criteria model.SearchCriteria) ([]model.Car, error)
	BookCar(booking model.Booking) error
}

type CarServiceImpl struct {
	carRepository repository.CarRepository
}

func NewCarService(carRepo repository.CarRepository) CarService {
	return &CarServiceImpl{carRepository: carRepo}
}

func (s *CarServiceImpl) SearchCars(criteria model.SearchCriteria) ([]model.Car, error) {
	return s.carRepository.FindCars(criteria)
}

func (s *CarServiceImpl) BookCar(booking model.Booking) error {
	// Check for existing bookings in the given time range
	bookings, err := s.carRepository.GetBookings(booking.CarID, booking.StartTime, booking.EndTime)
	if err != nil {
		return err
	}
	if len(bookings) > 0 {
		return errors.New("car is already booked for the selected time slot")
	}

	// Create new booking
	booking.Status = "confirmed"
	return s.carRepository.CreateBooking(booking)
}

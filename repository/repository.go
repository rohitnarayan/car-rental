package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/rohitnarayan/car-rental/model"
)

type CarRepository interface {
	FindCars(criteria model.SearchCriteria) ([]model.Car, error)
	CreateBooking(booking model.Booking) error
	GetBookings(carID uint, startTime, endTime time.Time) ([]model.Booking, error)
}

type CarRepositoryImpl struct {
	Cars     map[uint]model.Car
	bookings map[uint][]model.Booking
	mu       sync.Mutex
	nextID   uint
}

func NewCarRepository() CarRepository {
	return &CarRepositoryImpl{
		Cars:     make(map[uint]model.Car),
		bookings: make(map[uint][]model.Booking),
		nextID:   1,
	}
}

func (r *CarRepositoryImpl) FindCars(criteria model.SearchCriteria) ([]model.Car, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []model.Car
	for _, car := range r.Cars {
		if car.Status != "available" {
			continue
		}
		if criteria.Make != "" && car.Make != criteria.Make {
			continue
		}
		if criteria.Model != "" && car.Model != criteria.Model {
			continue
		}
		if criteria.MinPrice > 0 && car.Price < criteria.MinPrice {
			continue
		}
		if criteria.MaxPrice > 0 && car.Price > criteria.MaxPrice {
			continue
		}

		if !criteria.StartTime.IsZero() && !criteria.EndTime.IsZero() {
			if r.isCarBooked(car.ID, criteria.StartTime, criteria.EndTime) {
				continue
			}
		}

		result = append(result, car)
	}
	return result, nil
}

func (r *CarRepositoryImpl) CreateBooking(booking model.Booking) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.isCarBooked(booking.CarID, booking.StartTime, booking.EndTime) {
		return errors.New("car is already booked for the selected time slot")
	}

	booking.ID = r.nextID
	r.nextID++
	r.bookings[booking.CarID] = append(r.bookings[booking.CarID], booking)

	return nil
}

func (r *CarRepositoryImpl) GetBookings(carID uint, startTime, endTime time.Time) ([]model.Booking, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []model.Booking
	for _, booking := range r.bookings[carID] {
		if startTime.Before(booking.EndTime) && endTime.After(booking.StartTime) {
			result = append(result, booking)
		}
	}
	return result, nil
}

func (r *CarRepositoryImpl) isCarBooked(carID uint, startTime, endTime time.Time) bool {
	for _, booking := range r.bookings[carID] {
		if startTime.Before(booking.EndTime) && endTime.After(booking.StartTime) {
			return true
		}
	}
	return false
}

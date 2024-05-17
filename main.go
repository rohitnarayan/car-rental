package main

import (
	"fmt"
	"time"

	"github.com/rohitnarayan/car-rental/model"
	"github.com/rohitnarayan/car-rental/repository"
	"github.com/rohitnarayan/car-rental/service"
)

func main() {
	carRepo := repository.NewCarRepository()
	carService := service.NewCarService(carRepo)

	// Add some cars to the repository
	carRepo.(*repository.CarRepositoryImpl).Cars[1] = model.Car{ID: 1, Make: "Toyota", Model: "Corolla", Year: 2020, Price: 50, Status: "available"}
	carRepo.(*repository.CarRepositoryImpl).Cars[2] = model.Car{ID: 2, Make: "Honda", Model: "Civic", Year: 2021, Price: 60, Status: "available"}
	carRepo.(*repository.CarRepositoryImpl).Cars[3] = model.Car{ID: 3, Make: "Ford", Model: "Mustang", Year: 2022, Price: 120, Status: "available"}

	// Create a booking for testing
	startTime := time.Now().Add(24 * time.Hour)
	endTime := startTime.Add(48 * time.Hour)
	booking := model.Booking{CarID: 1, UserID: 1, StartTime: startTime, EndTime: endTime, Status: "confirmed"}
	err := carService.BookCar(booking)
	if err != nil {
		fmt.Println("Error booking car:", err)
	}

	// Search for available cars in a given time range
	criteria := model.SearchCriteria{
		StartTime: startTime.Add(12 * time.Hour),
		EndTime:   endTime.Add(24 * time.Hour),
		MinPrice:  40,
		MaxPrice:  130,
	}

	cars, err := carService.SearchCars(criteria)
	if err != nil {
		fmt.Println("Error searching cars:", err)
	} else {
		fmt.Println("Available cars from", criteria.StartTime, "to", criteria.EndTime, ":")
		for _, car := range cars {
			fmt.Printf("ID: %d, Make: %s, Model: %s, Price: %.2f\n", car.ID, car.Make, car.Model, car.Price)
		}
	}
}

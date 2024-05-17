package model

import "time"

type Car struct {
	ID     uint    `json:"id"`
	Make   string  `json:"make"`
	Model  string  `json:"model"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

type Booking struct {
	ID        uint      `json:"id"`
	CarID     uint      `json:"car_id"`
	UserID    uint      `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
}

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

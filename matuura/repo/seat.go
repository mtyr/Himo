package repo

import "github.com/mtyr/Himo/matuura/model"

var seats []model.Seat

func SetSeat(seat model.Seat) {
	seats = append(seats, seat)
}

func GetAllSeats() []model.Seat {
	return seats
}

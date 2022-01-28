package repository

import "github.com/YanAmorelli/bookings/internal/models"

type DatabaseRepo interface {
	InsertReservations(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
}
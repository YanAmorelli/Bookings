package models

import "time"

type Reservation struct {
	ID			int
	FirstName 	string
	LastName  	string
	Email     	string
	Phone     	string
	StartDate	time.Time
	EndDate		time.Time
	RoomId 		int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Room 		Room
}

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Room struct {
	ID		 	int
	RoomName	string
	CreatedAt   time.Time
	UpdatedAt   time.Time	
}

type Restriction struct {
	ID		 		int
	RestrictionName	string
	CreatedAt   	time.Time
	UpdatedAt   	time.Time	
}

type RoomRestriction struct {
	ID		 		int
	StartDate		time.Time
	EndDate			time.Time
	RoomId 			int
	ReservationId 	int
	CreatedAt   	time.Time
	UpdatedAt   	time.Time	
	RestrictionId	int
	Room 			Room
	Reservation 	Reservation
	Restriction		Restriction
}
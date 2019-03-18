package model

type NotificationType int32
type NotificationMethod int32

const (
	ALERT   NotificationType = 0
	INFORM  NotificationType = 1
	EMAIL NotificationMethod = 10
	SMS NotificationMethod = 11
)



type Booking struct {
	ID     int64
	ShowingId    int64
	Amount	int64
	PhoneNumber string
	OrderId string
}

type Notification struct {
	Type NotificationType
	Method NotificationMethod
	Payload Payload
}

type Payload struct {
	Cinema string
	Movie string
	AvailableSeats int64
	TotalSeats int64
	BookingUser User
}

type User struct {
	Name string
	LastName string
	Email string
	PhoneNumber string
}



package parking

import "time"

type CheckInFailEvent struct {
	Plate     string
	LocalTime time.Time
}

type CheckOutFailEvent struct {
	Plate     string
	LocalTime time.Time
}

type CheckInEvent struct {
	Plate     string
	LocalTime time.Time
}

type CheckOutEvent struct {
	Plate     string
	LocalTime time.Time
}

type PaidEvent struct {
	Plate   string
	PayTime time.Time
	Amount  int
}

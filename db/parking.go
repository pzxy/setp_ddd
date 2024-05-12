package db

import (
	"time"
)

type Parking struct {
	ID           string //plate
	CheckInTime  time.Time
	LastPlayTime time.Time
	TotalPaid    int
	CreatedAt    time.Time `gorm:"<-:create"`
	UpdatedAt    time.Time
}

func (*Parking) TableName() string {
	return "parking_table"
}

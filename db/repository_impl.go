package db

import (
	"github.com/step_ddd/domain/parking"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var DB = &RepositoryImpl{}

type RepositoryImpl struct {
	db *gorm.DB
}

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.db = db

	migrator := db.Migrator()
	if err = migrator.AutoMigrate(&Parking{}); err != nil {
		panic(err)
	}
}

func (r *RepositoryImpl) FindById(plate string) (parking.Parking, error) {
	p := &Parking{ID: plate}
	r.db.FirstOrCreate(p)
	return &parking.ParkingImpl{
		Id:           p.ID,
		CheckInTime:  p.CheckInTime,
		LastPlayTime: p.LastPlayTime,
		TotalPaid:    p.TotalPaid,
	}, nil
}

func (r *RepositoryImpl) SaveCheckIn(command *parking.CheckInCommand) error {
	r.db.Model(&Parking{ID: command.Plate}).Updates(map[string]interface{}{
		"check_in_time": time.Now(),
	})
	return nil
}

func (r *RepositoryImpl) SaveCheckOut(command *parking.CheckOutCommand) error {
	//TODO implement me
	panic("implement me")
}

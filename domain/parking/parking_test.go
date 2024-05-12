package parking

import (
	"testing"
	"time"
)

func TestCalcFeeNow(t *testing.T) {
	p := &ParkingImpl{
		Id:           "浙：JD12345",
		CheckInTime:  time.Date(2024, 5, 10, 19, 30, 0, 0, time.UTC),
		LastPlayTime: time.Time{},
	}
	fee, err := p.calcFeeNow(time.Now())
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(fee)
}

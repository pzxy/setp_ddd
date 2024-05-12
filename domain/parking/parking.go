package parking

import (
	"fmt"
	"github.com/step_ddd/log"
	"time"
)

// Parking 停车聚合对象
type Parking interface {
	handleCheckIn(EventQueue, *CheckInCommand) error
	calcFeeNow(time.Time) (int, error)
	handlePay(queue EventQueue, command *NotifyPayCommand) bool
	handleCheckOut(queue EventQueue, command *CheckOutCommand) bool
}

type ParkingImpl struct {
	Id           string
	CheckInTime  time.Time
	LastPlayTime time.Time
	TotalPaid    int
}

func (p *ParkingImpl) handleCheckIn(event EventQueue, command *CheckInCommand) error {
	if p.inPark() {
		event.Enqueue(&CheckInFailEvent{p.Id, time.Now()})
		err := fmt.Errorf("resource:%s already exists", p.Id)
		log.G.Error(err.Error())
		return err
	}
	event.Enqueue(&CheckInEvent{p.Id, time.Now()})
	return nil
}

func (p *ParkingImpl) handlePay(event EventQueue, command *NotifyPayCommand) bool {
	if p.inPark() {
		return false
	}
	p.LastPlayTime = command.payTime
	p.TotalPaid += command.amount
	event.Enqueue(PaidEvent{p.Id, p.LastPlayTime, command.amount})
	return true
}
func (p *ParkingImpl) handleCheckOut(event EventQueue, command *CheckOutCommand) bool {
	if p.inPark() {
		event.Enqueue(&CheckOutFailEvent{p.Id, time.Now()})
		return false
	}
	fee, err := p.calcFeeNow(time.Now())
	if err != nil || fee > 0 {
		return false
	}
	p.CheckInTime = time.Time{}
	p.TotalPaid = 0
	p.LastPlayTime = time.Time{}
	event.Enqueue(&CheckOutEvent{Plate: p.Id, LocalTime: time.Now()})
	return true
}

func (p *ParkingImpl) inPark() bool {
	return !p.CheckInTime.Equal(time.Time{})
}

func (p *ParkingImpl) calcFeeNow(now time.Time) (int, error) {
	if p.CheckInTime.Equal(time.Time{}) {
		return -1, fmt.Errorf("车辆 %s 没有入场", p.Id)
	}
	if p.LastPlayTime.Equal(time.Time{}) {
		return timeBetween(p.CheckInTime, now), nil
	}
	return 1, nil
}

func timeBetween(start time.Time, end time.Time) int {
	a := start.Sub(end).Hours()
	return int(a)
}

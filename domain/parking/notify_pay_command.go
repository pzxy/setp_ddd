package parking

import "time"

type NotifyPayCommand struct {
	plate   string
	amount  int
	payTime time.Time
}

type NotifyPayCommandHandler struct {
	Repository
}

func (c *NotifyPayCommandHandler) handle(command *CalcFeeCommand) (int, error) {
	parking, _ := c.Repository.FindById(command.plate)
	return parking.calcFeeNow(time.Now())
}

package parking

import "time"

type CalcFeeCommand struct {
	plate string
}

type CalcFeeCommandHandler struct {
	Repository
}

func (c *CalcFeeCommandHandler) handle(command *CalcFeeCommand) (int, error) {
	parking, _ := c.Repository.FindById(command.plate)
	return parking.calcFeeNow(time.Now())
}

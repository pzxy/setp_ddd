package parking

type CheckInCommand struct {
	Plate string
}

type CheckInCommandHandler struct {
	Repository
}

func (c *CheckInCommandHandler) Handle(event EventQueue, command *CheckInCommand) error {
	p, err := c.FindById(command.Plate)
	if err != nil {
		return err
	}
	if err = p.handleCheckIn(event, command); err != nil {
		return err
	}
	return c.SaveCheckIn(command)
}

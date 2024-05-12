package parking

type CheckOutCommand struct {
	plate string
}

type CheckOutCommandHandler struct {
	Repository
}

func (c *CheckOutCommandHandler) handle(event EventQueue, command *CheckOutCommand) {
	p, _ := c.Repository.FindById(command.plate)
	p.handleCheckOut(event, command)
	_ = c.Repository.SaveCheckOut(command)
}

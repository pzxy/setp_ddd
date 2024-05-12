package parking

type Repository interface {
	FindById(plate string) (Parking, error)
	SaveCheckIn(*CheckInCommand) error
	SaveCheckOut(*CheckOutCommand) error
}

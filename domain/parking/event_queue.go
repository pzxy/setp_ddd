package parking

type EventQueue interface {
	Enqueue(interface{})
}

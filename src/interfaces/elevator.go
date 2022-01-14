package behaviour

type ElevatorBehavior interface {
	Advance(targetFloor []int) error
}

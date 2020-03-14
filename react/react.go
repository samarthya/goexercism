package react

//IValue interface
type IValue interface {
	CreateInput() InputCell
}

// InputCell cell value
type InputCell struct {
	value int
}

// SetValue On the input value
func (i *InputCell) SetValue(val int) {
	i.value = val
}

//Value evaluates value
func (i *InputCell) Value() int {
	return i.value
}

// Reactor type reactor
type Reactor struct {
	ic InputCell
}

//CreateInput Creates input cell
func (r *Reactor) CreateInput(i int) {
	// Sets the value
	r.ic.SetValue(i)
}

// New function
func New() Reactor {
	return &Reactor{}
}

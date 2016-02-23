package react

// CreateInput creates an input cell linked into the reactor with the given
// initial value.
func (r *Reactor) CreateInput(int) InputCell {
}

// CreateCompute1 creates a compute cell which computes its value based on
// one other cell. The compute function will only be called if the value of
// the passed cell changes.
func (r *Reactor) CreateCompute1(Cell, func(int) int) ComputeCell {
}

	// CreateCompute2 is like CreateCompute1, but depending on two cells.
	// The compute function will only be called if the value of any of the
	// passed cells changes.
	CreateCompute2(Cell, Cell, func(int, int) int) ComputeCell
}

// Value returns the current value of the cell.
func (c *Cell) Value() int {
}

// SetValue sets the value of the cell.
func (ic *InputCell) SetValue(int) {
}

// A ComputeCell always computes its value based on other cells and can
// call callbacks upon changes.
type ComputeCell interface {
	Cell

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (cc *ComputeCell) AddCallback(func(int)) CallbackHandle {
}

// RemoveCallback removes a previously added callback, if it exists.
func (cc *ComputeCell) RemoveCallback(CallbackHandle) {
}

package react

const testVersion = 4

type Body struct {}
type BCell struct {
}
type Value struct {
	BCell
	value int
}
type Computer struct {
	BCell
	in1, in2 Value
	value int
	callbacks []func(int)
}

func New() (r *Body) {
	return new(Body)
}

// CreateInput creates an input cell linked into the reactor with the given
// initial value.
func (r *Body) CreateInput(v int) Value {
	return Value{value: v}
}

// CreateCompute1 creates a compute cell which computes its value based on
// one other cell. The compute function will only be called if the value of
// the passed cell changes.
func (r *Body) CreateCompute1(v Cell, f func(int) int) Computer {
	c = Computer{in1: v, value: f(v), callbacks: []func(int){f}}
	// FIXME
	valueCallback = func(newVal Value) {
		observed = append(observed, v)
	}
}

// CreateCompute2 is like CreateCompute1, but depending on two cells. The
// compute function will only be called if the value of any of the passed
// cells changes.
func (r *Body) CreateCompute2(v1 Value, v2 Value, func(int, int) int) Computer {
	return Computer{in1: v1, in2: v2, value: f(v1, v2), callbacks: []func(int){f}}
}

// Value returns the current value of the cell.
func (c *Value) Value() int {
	return c.value
}

// Value returns the current value of the cell.
func (c *Computer) Value() int {
	return c.value
}

// SetValue sets the value of the cell.
func (val *Value) SetValue(v int) {
	val.value = v
}

// SetValue sets the value of the cell.
func (val *Computer) SetValue(v int) {
	val.value = v
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (c *Computer) AddCallback(f func(int)) CallbackHandle {
	c.callbacks = append(c.callbacks, f)
	return f
}

// RemoveCallback removes a previously added callback, if it exists.
func (c *Computer) RemoveCallback(f CallbackHandle) {
	delete(c.callbacks, f)
}

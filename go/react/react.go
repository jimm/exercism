package react

const testVersion = 4

type JReactor struct{}
type JCell struct {
	value     int
	observers []*JCell
	inputs    []*JCell
	f         func([]*JCell) int
	callbacks []func(int)
}
type JCallbackHandle struct {
	index int
}

func New() Reactor {
	return new(JReactor)
}

// CreateInput creates an input cell linked into the reactor with the given
// initial value.
func (r *JReactor) CreateInput(v int) InputCell {
	return &JCell{value: v}
}

// CreateCompute1 creates a compute cell which computes its value based on
// one other cell. The compute function will only be called if the value of
// the passed cell changes.
func (r *JReactor) CreateCompute1(v Cell, f func(int) int) ComputeCell {
	in := v.(*JCell)
	wrapper := func(args []*JCell) int {
		return f(args[0].Value())
	}
	c := &JCell{value: f(in.Value()), f: wrapper, inputs: []*JCell{in}}
	in.addObserver(c)
	return c
}

// CreateCompute2 is like CreateCompute1, but depending on two cells. The
// compute function will only be called if the value of any of the passed
// cells changes.
func (r *JReactor) CreateCompute2(v1 Cell, v2 Cell, f func(int, int) int) ComputeCell {
	in1 := v1.(*JCell)
	in2 := v2.(*JCell)
	wrapper := func(args []*JCell) int {
		return f(args[0].Value(), args[1].Value())
	}
	c := &JCell{value: f(in1.Value(), in2.Value()), f: wrapper, inputs: []*JCell{in1, in2}}
	in1.addObserver(c)
	in2.addObserver(c)
	return c
}

// Value returns the current value of the cell.
func (c JCell) Value() int {
	return c.value
}

// SetValue sets the value of the cell.
func (c *JCell) SetValue(v int) {
	c.value = v
	c.changed()
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (c *JCell) AddCallback(f func(int)) CallbackHandle {
	c.callbacks = append(c.callbacks, f)
	return JCallbackHandle{len(c.callbacks) - 1}
}

// RemoveCallback removes a previously added callback, if it exists.
func (c *JCell) RemoveCallback(cbh CallbackHandle) {
	indexer := cbh.(JCallbackHandle)
	newCallbacks := make([]func(int), 0)
	for i, f := range c.callbacks {
		if i != indexer.index {
			newCallbacks = append(newCallbacks, f)
		}
	}
	c.callbacks = newCallbacks
}

func (v *JCell) addObserver(c ComputeCell) {
	jc := c.(*JCell)
	v.observers = append(v.observers, jc)
}

func (v *JCell) changed() {
	for _, o := range v.observers {
		o.value = o.f(o.inputs)
		for _, f := range o.callbacks {
			f(o.value)
		}
	}
}

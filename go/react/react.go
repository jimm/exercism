package react

const testVersion = 4

type JReactor struct{}
type JCell struct {
}
type JInputCell struct {
	JCell
	value     int
	observers []ComputeCell
}
type CPU struct {
	in1, in2  JInputCell
	value     int
	callbacks []func(int)
}
type JComputeCell struct {
	JCell
}
type JComputeCell1 struct {
	JComputeCell
	cpu CPU
	f   func(int) int
}
type JComputeCell2 struct {
	JComputeCell
	cpu CPU
	f   func(int, int) int
}
type JCallbackHandle struct {
	index int
}

func New() Reactor {
	return new(JReactor)
}

// CreateInput creates an input cell linked into the reactor with the given
// initial value.
func (r JReactor) CreateInput(v int) InputCell {
	return JInputCell{value: v}
}

// CreateCompute1 creates a compute cell which computes its value based on
// one other cell. The compute function will only be called if the value of
// the passed cell changes.
func (r JReactor) CreateCompute1(v Cell, f func(int) int) ComputeCell {
	in := v.(JInputCell)
	cpu := CPU{in1: in, value: f(in.Value())}
	c := JComputeCell1{cpu: cpu, f: f}
	in.addObserver(c)
	return c
}

// CreateCompute2 is like CreateCompute1, but depending on two cells. The
// compute function will only be called if the value of any of the passed
// cells changes.
func (r JReactor) CreateCompute2(v1 Cell, v2 Cell, f func(int, int) int) ComputeCell {
	in1 := v1.(JInputCell)
	in2 := v2.(JInputCell)
	cpu := CPU{in1: in1, in2: in2, value: f(in1.Value(), in2.Value())}
	c := JComputeCell2{cpu: cpu, f: f}
	in1.addObserver(c)
	in2.addObserver(c)
	return c
}

// Value returns the current value of the cell.
func (c JInputCell) Value() int {
	return c.value
}

// Value returns the current value of the cell.
func (c JComputeCell1) Value() int {
	return c.cpu.value
}

// Value returns the current value of the cell.
func (c JComputeCell2) Value() int {
	return c.cpu.value
}

// SetValue sets the value of the cell.
func (val JInputCell) SetValue(v int) {
	val.value = v
	val.changed()
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (c JComputeCell1) AddCallback(f func(int)) CallbackHandle {
	return c.cpu.addCallback(f)
}
func (c JComputeCell2) AddCallback(f func(int)) CallbackHandle {
	return c.cpu.addCallback(f)
}

func (cpu CPU) addCallback(f func(int)) CallbackHandle {
	cpu.callbacks = append(cpu.callbacks, f)
	return JCallbackHandle{len(cpu.callbacks) - 1}
}

// RemoveCallback removes a previously added callback, if it exists.
func (c JComputeCell1) RemoveCallback(cbh CallbackHandle) {
	c.cpu.removeCallback(cbh)
}
func (c JComputeCell2) RemoveCallback(cbh CallbackHandle) {
	c.cpu.removeCallback(cbh)
}

func (cpu CPU) removeCallback(cbh CallbackHandle) {
	indexer := cbh.(JCallbackHandle)
	newCallbacks := make([]func(int), 0)
	// newCallbacks = append(newCallbacks, callbacks[0:index])
	// newCallbacks = append(newCallbacks, callbacks[index+1])
	for i, f := range cpu.callbacks {
		if i != indexer.index {
			newCallbacks = append(newCallbacks, f)
		}
	}
	cpu.callbacks = newCallbacks
}

func (v JInputCell) addObserver(c ComputeCell) {
	v.observers = append(v.observers, c)
}

func (v JInputCell) changed() {
	for _, o := range v.observers {
		c1, ok := o.(JComputeCell1)
		if ok {
			c1.changed(v)
		} else {
			c2 := o.(JComputeCell2)
			c2.changed(v)
		}
	}
}

func (c JComputeCell1) changed(v JInputCell) {
	cpu := c.cpu
	cpu.value = c.f(cpu.in1.Value())
	for _, f := range cpu.callbacks {
		f(cpu.value)
	}
}

func (c JComputeCell2) changed(v JInputCell) {
	cpu := c.cpu
	cpu.value = c.f(cpu.in1.Value(), cpu.in2.Value())
	for _, f := range cpu.callbacks {
		f(cpu.value)
	}
}

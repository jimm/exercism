package react

const testVersion = 4

var nextCallbackId = 0

type JReactor struct {
	working bool
	changed map[*JCell]bool
}
type JCell struct {
	reactor   *JReactor
	value     int
	OrigValue interface{}
	observers []*JCell
	inputs    []*JCell
	f         func([]*JCell) int
	callbacks map[int]func(int)
}
type JCallbackHandle struct {
	id int
}

// **************** reactor ****************

func New() Reactor {
	return &JReactor{working: false}
}

// CreateInput creates an input cell linked into the reactor with the given
// initial value.
func (r *JReactor) CreateInput(v int) InputCell {
	return &JCell{reactor: r, value: v}
}

// CreateCompute1 creates a compute cell which computes its value based on
// one other cell. The compute function will only be called if the value of
// the passed cell changes.
func (r *JReactor) CreateCompute1(v Cell, f func(int) int) ComputeCell {
	in := v.(*JCell)
	wrapper := func(args []*JCell) int {
		return f(args[0].Value())
	}
	c := &JCell{reactor: r, value: f(in.Value()), f: wrapper,
		inputs: []*JCell{in}, callbacks: make(map[int]func(int), 0)}
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
	c := &JCell{reactor: r, value: f(in1.Value(), in2.Value()), f: wrapper,
		inputs: []*JCell{in1, in2}, callbacks: make(map[int]func(int), 0)}
	in1.addObserver(c)
	in2.addObserver(c)
	return c
}

func (r *JReactor) startWorking() {
	r.working = true
	r.changed = map[*JCell]bool{}
}

func (r *JReactor) addChangedCell(c *JCell) {
	r.changed[c] = true
}

func (r *JReactor) stopWorking() {
	for c := range r.changed {
		val := c.Value()
		if val != c.OrigValue {
			for _, f := range c.callbacks {
				f(val)
			}
		}
		c.OrigValue = nil
	}
	r.working = false
}

// **************** cell ****************

// Value returns the current value of the cell.
func (c *JCell) Value() int {
	return c.value
}

// SetValue sets the value of the cell. This isn't thread safe, but then
// again the tests don't require that it be thread safe.
func (c *JCell) SetValue(v int) {
	r := c.reactor
	wasWorking := r.working
	if !wasWorking {
		r.startWorking()
	}
	if c.value != v {
		c.saveValue(v)
		r.addChangedCell(c)
	}
	if !wasWorking {
		r.stopWorking()
	}
}

// saveValue sets the new value of the cell and notifies observers. If the
// cell has not yet saved its original value, then do so.
func (c *JCell) saveValue(v int) {
	if c.OrigValue == nil {
		c.OrigValue = c.value
	}
	c.value = v
	c.notifyObservers()
}

func (c *JCell) notifyObservers() {
	for _, o := range c.observers {
		o.SetValue(o.f(o.inputs)) // recalculate
	}
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (c *JCell) AddCallback(f func(int)) CallbackHandle {
	c.callbacks[nextCallbackId] = f
	nextCallbackId++
	return JCallbackHandle{id: nextCallbackId - 1}
}

// RemoveCallback removes a previously added callback, if it exists.
func (c *JCell) RemoveCallback(cbh CallbackHandle) {
	jcbh := cbh.(JCallbackHandle)
	delete(c.callbacks, jcbh.id)
}

func (c *JCell) addObserver(cc ComputeCell) {
	jc := cc.(*JCell)
	c.observers = append(c.observers, jc)
}

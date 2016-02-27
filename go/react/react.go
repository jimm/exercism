package react

const testVersion = 4

var nextCallbackId = 0

type reactor struct {
	workLevel int
	changed   map[*cell]int
}
type cell struct {
	reactor   *reactor
	value     int
	origValue interface{} // nil if no change to cell
	observers []*cell
	inputs    []*cell
	f         func([]*cell) int
	callbacks map[int]func(int)
}
type callbackHandle struct {
	id int
}

// **************** reactor ****************

func New() Reactor {
	return &reactor{}
}

// CreateInput creates an input cell linked into the reactor with the given
// initial value.
func (r *reactor) CreateInput(v int) InputCell {
	return &cell{reactor: r, value: v}
}

// CreateCompute1 creates a compute cell which computes its value based on
// one other cell. The compute function will only be called if the value of
// the passed cell changes.
func (r *reactor) CreateCompute1(v Cell, f func(int) int) ComputeCell {
	in := v.(*cell)
	wrapper := func(args []*cell) int {
		return f(args[0].Value())
	}
	c := &cell{reactor: r, value: f(in.Value()), f: wrapper,
		inputs: []*cell{in}, callbacks: make(map[int]func(int), 0)}
	in.addObserver(c)
	return c
}

// CreateCompute2 is like CreateCompute1, but depending on two cells. The
// compute function will only be called if the value of any of the passed
// cells changes.
func (r *reactor) CreateCompute2(v1 Cell, v2 Cell, f func(int, int) int) ComputeCell {
	in1 := v1.(*cell)
	in2 := v2.(*cell)
	wrapper := func(args []*cell) int {
		return f(args[0].Value(), args[1].Value())
	}
	c := &cell{reactor: r, value: f(in1.Value(), in2.Value()), f: wrapper,
		inputs: []*cell{in1, in2}, callbacks: make(map[int]func(int), 0)}
	in1.addObserver(c)
	in2.addObserver(c)
	return c
}

func (r *reactor) startWorking() {
	if r.workLevel == 0 {
		r.changed = map[*cell]int{}
	}
	r.workLevel++
}

// addChangingCell remembers cells that are about to be recalculated and
// saves their initial values.
func (r *reactor) addChangingCell(c *cell) {
	_, found := r.changed[c]
	if !found {
		r.changed[c] = c.Value()
	}
}

func (r *reactor) stopWorking() {
	r.workLevel--
	if r.workLevel != 0 {
		return
	}

	for c, origValue := range r.changed {
		val := c.Value()
		if val != origValue {
			for _, f := range c.callbacks {
				f(val)
			}
		}
		c.origValue = nil
	}
}

// **************** cell ****************

// Value returns the current value of the cell.
func (c *cell) Value() int {
	return c.value
}

// SetValue sets the value of the cell. This isn't thread safe, but then
// again the tests don't require that it be thread safe.
func (c *cell) SetValue(v int) {
	r := c.reactor
	r.startWorking()
	if c.value != v {
		r.addChangingCell(c)
		c.saveValue(v)
	}
	r.stopWorking()
}

// saveValue sets the new value of the cell and notifies observers. If the
// cell has not yet saved its original value, then do so.
func (c *cell) saveValue(v int) {
	c.value = v
	c.notifyObservers()
}

func (c *cell) notifyObservers() {
	for _, o := range c.observers {
		o.SetValue(o.f(o.inputs)) // recalculate
	}
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a callback handle which can be used to remove the callback.
func (c *cell) AddCallback(f func(int)) CallbackHandle {
	c.callbacks[nextCallbackId] = f
	nextCallbackId++
	return callbackHandle{id: nextCallbackId - 1}
}

// RemoveCallback removes a previously added callback, if it exists.
func (c *cell) RemoveCallback(cbh CallbackHandle) {
	cbhStruct := cbh.(callbackHandle)
	delete(c.callbacks, cbhStruct.id)
}

func (c *cell) addObserver(cc ComputeCell) {
	cStruct := cc.(*cell)
	c.observers = append(c.observers, cStruct)
}

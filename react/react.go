package react

type computeCell struct {
	value     int
	deps      []Cell
	valueFunc interface{}
	cid       int
	callbacks map[int]func(int)
}

type canceler struct {
	cid int
	cc  *computeCell
}

type inputCell struct {
	r     *reactor
	value int
}

type reactor struct {
	icells []*inputCell
	ccells []*computeCell
}

func (r *reactor) CreateInput(value int) InputCell {
	c := &inputCell{value: value, r: r}
	r.icells = append(r.icells, c)
	return c
}

func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	cc := &computeCell{value: f(c.Value()), valueFunc: f, deps: []Cell{c}}
	r.ccells = append(r.ccells, cc)
	return cc
}

func (r *reactor) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	cc := &computeCell{value: f(c1.Value(), c2.Value()), valueFunc: f, deps: []Cell{c1, c2}}
	r.ccells = append(r.ccells, cc)
	return cc
}

func (c *inputCell) Value() int {
	return c.value
}

func (c *inputCell) SetValue(value int) {
	changed := false
	if c.value != value {
		changed = true
	}
	c.value = value
	if changed {
		changedCells := map[Cell]bool{c: true}
		for _, cc := range c.r.ccells {
			depChange := false
			for _, dep := range cc.deps {
				if changedCells[dep] {
					depChange = true
					break
				}
			}
			if depChange {
				v := cc.calc()
				if v != cc.value {
					cc.value = v
					for _, cb := range cc.callbacks {
						cb(v)
					}
					changedCells[cc] = true
				}
			}
		}
	}
}

func (c *computeCell) calc() int {
	v := 0
	switch f := c.valueFunc.(type) {
	case func(int) int:
		v = f(c.deps[0].Value())
	case func(int, int) int:
		v = f(c.deps[0].Value(), c.deps[1].Value())
	}
	return v
}
func (c *computeCell) Value() int {
	return c.value
}

func (c *computeCell) AddCallback(cf func(int)) Canceler {
	if c.cid == 0 {
		c.callbacks = make(map[int]func(int))
	}
	c.cid++
	c.callbacks[c.cid] = cf
	return &canceler{cid: c.cid, cc: c}
}

func (c *canceler) Cancel() {
	delete(c.cc.callbacks, c.cid)
}

// New a reactor
func New() Reactor {
	return &reactor{}
}

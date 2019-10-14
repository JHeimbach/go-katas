package kata

var actionFuncs = []func(int) int{
	func(x int) int { return x + 1 },
	func(x int) int { return 0 },
	func(x int) int { return x / 2 },
	func(x int) int { return x * 100 },
	func(x int) int { return x % 2 },
}

func Actions() []func(int) int {
	return actionFuncs
}

type Machine struct {
	cmd     int
	mapping map[int]int
}

//Function called to get your machine initialised
func NewMachine() Machine {
	return Machine{mapping: make(map[int]int)}
}

func (m *Machine) Command(cmd int, num int) int {
	m.cmd = cmd
	return Actions()[m.mapping[cmd]](num)
}

func (m *Machine) Response(res bool) {
	if !res {
		m.mapping[m.cmd] = (m.mapping[m.cmd] + 1) % len(Actions())
	}
}

package kata

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewMachine(t *testing.T) {
	_actions := []func(int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return 0 },
		func(x int) int { return x / 2 },
		func(x int) int { return x * 100 },
		func(x int) int { return x % 2 }}

	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	t.Run("Should apply the num * 0 action to the command 0", func(t *testing.T) {
		mach := NewMachine()

		for i := 0; i < 20; i++ {
			number := mach.Command(0, rnd.Intn(10))
			mach.Response(number == 0)
		}
		got := mach.Command(0, 1000)
		if got != 0 {
			t.Errorf("got %d want %d", got, 0)
		}
	})

	t.Run("train and evaluate", func(t *testing.T) {
		mach := NewMachine()
		for i := 0; i < 200; i++ {
			number := rnd.Intn(100)
			num := mach.Command(i%5, number)
			mach.Response(_actions[i%5](number) == num)
		}

		type test struct {
			action int
			num    int
			want   int
			name   string
		}

		tests := []test{
			{num: 100, want: 101, name: "Should apply the num + 1 action to the command 0"},
			{action: 1, num: 100, name: "Should apply the num * 0 action to the command 1"},
			{action: 2, num: 100, want: 50, name: "Should apply the num / 2 action to the command 2"},
			{action: 3, num: 1, want: 100, name: "Should apply the num * 100 action to the command 3"},
			{action: 4, num: 100, name: "Should apply the num % 2 action to the command 4"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := mach.Command(tt.action, tt.num)
				if got != tt.want {
					t.Errorf("got %d want %d", got, tt.want)
				}
			})
		}
	})
}

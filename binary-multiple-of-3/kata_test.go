package kata

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func TestMultipleOfThreeRegex(t *testing.T) {
	type test struct {
		input string
		want  bool
	}
	simpleTests := []test{
		{" 0", false},
		{"abc", false},
		{"000", true},
		{"110", true},
		{"111", false},
		{strconv.FormatInt(12345678, 2), true},
	}
	for _, tt := range simpleTests {
		t.Run(fmt.Sprintf("should test that the solution returns the correct value for %q", tt.input), func(t *testing.T) {
			got, _ := regexp.Match(MultipleOf3Regex, []byte(tt.input))
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

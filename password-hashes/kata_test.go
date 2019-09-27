package kata

import "testing"

func TestPassHash(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "solution returns the correct value",
			in:   "password",
			out:  "5f4dcc3b5aa765d61d8327deb882cf99",
		},
		{
			name: "hash string must be 32",
			in:   "abc123",
			out:  "e99a18c428cb38d5f260853678922e03",
		},
		{
			name: "Should handle empty string",
			in:   "",
			out:  "d41d8cd98f00b204e9800998ecf8427e",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PassHash(tt.in)

			if got != tt.out {
				t.Errorf("PassHash(%q) returned %q, expected %q", tt.in, got, tt.out)
			}
		})
	}
}

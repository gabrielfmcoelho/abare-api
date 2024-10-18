package mock

import "testing"

func Print() string {
	return "Hello, World!"
}

func TestMock(t *testing.T) {
	want := "Hello, World!"
	if got := Print(); got != want {
		t.Errorf("Print() = %q, want %q", got, want)
	}
}

package components

import (
	"testing"
)

func TestPlayer(t *testing.T) {
	var list = []struct {
		name     string
		mark     string
		expected Player
	}{
		{"Ameya", "X", Player{Mark: "X", Name: "Ameya"}},
		{"Meghan", "O", Player{Mark: "O", Name: "Meghan"}},
		{"Rishab", "X", Player{Mark: "X", Name: "Rishab"}},
		{"Devesh", "X", Player{Mark: "X", Name: "Devesh"}},
		{"Eshan", "O", Player{Mark: "O", Name: "Eshan"}},
	}
	for _, str := range list {
		actual := CreatePlayer(str.name, str.mark)
		if *actual != str.expected {
			t.Error("In CreatePlayer Actual is :", *actual, "But expected is : ", str.expected)
		}
	}
}

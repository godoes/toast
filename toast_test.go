package toast

import (
	"testing"
)

func TestToast(t *testing.T) {
	notification := Notification{
		AppID:   "Example App",
		Title:   "My notification",
		Message: "Some message about how important something is...",
		//Icon:    "go.png", // This file must exist (remove this line if it doesn't)
		Actions: []Action{
			{"protocol", "I'm a button", ""},
			{"protocol", "Me too!", ""},
		},
	}
	err := notification.Push()
	if err != nil {
		t.Error(err)
	}
}

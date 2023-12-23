package kmnbot_test

import (
	"testing"

	"github.com/kinabcd/kmnbot"
)

func TestHandbook(t *testing.T) {
	if len(kmnbot.Handbook) == 0 {
		t.Errorf("empty data")
	}
	for name, item := range kmnbot.Handbook {
		t.Logf("%s %v", name, item)
	}
}

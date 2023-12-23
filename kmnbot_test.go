package kmnbot_test

import (
	"encoding/json"
	"testing"

	"github.com/kinabcd/kmnbot"
)

func TestFetchBox(t *testing.T) {
	kmns, err := kmnbot.FetchBox("kinabcd")
	if err != nil {
		t.Error(err)
	}
	bytes, err := json.Marshal(kmns)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))
}

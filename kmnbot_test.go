package kmnbot_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/kinabcd/kmnbot"
)

func TestFetchBox(t *testing.T) {
	if os.Getenv("Github") != "" {
		t.Skip("Skip on Github")
	}
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

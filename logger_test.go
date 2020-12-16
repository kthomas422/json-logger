package logger

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	var dasLogger = NewLogger(os.Stdout)
	if err := dasLogger.Info(SimpleLog("hi")); err != nil {
		t.Errorf("error from logging: %v", err)
	}
	if err := dasLogger.Error(SimpleLog("poop")); err != nil {
		t.Errorf("error from logging: %v", err)
	}
}

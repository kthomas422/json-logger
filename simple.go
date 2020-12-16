package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Simple contains a log message in a string.
type Simple struct {
	Msg string
}

// SimpleLog creates a new struct the message
// (intended for myLogger.Info(logger.SimpleLog(msg))
func SimpleLog(msg string) Simple {
	return Simple{Msg: msg}
}

// WriteInfo satisfies the Log interface, it writes a json line to the log output
func (sl Simple) WriteInfo(w io.Writer) error {
	logBytes, err := sl.getLogBytes("info")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

// WriteError satisfies the Log interface, it writes a json line to the log output
func (sl Simple) WriteError(w io.Writer) error {
	logBytes, err := sl.getLogBytes("error")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

// getLogBytes turns the log message into a json string
func (sl *Simple) getLogBytes(lvl string) ([]byte, error) {
	log := map[string]string{
		"time":  time.Now().Format(time.RFC3339),
		"level": lvl,
		"msg":   sl.Msg,
	}
	return json.Marshal(log)
}

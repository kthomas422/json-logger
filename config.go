package jsonlogger

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// ConfigLog is a wrapper for your application's configuration struct to log
type ConfigLog struct {
	Config interface{}
	Errors []error
}

// Logs the configuration
func (cl ConfigLog) WriteInfo(w io.Writer) error {
	log := map[string]interface{}{
		"time":   time.Now().Format(time.RFC3339),
		"type":   "config",
		"level":  "info",
		"config": cl.Config,
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

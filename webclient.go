package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// ClientRequestLog wraps up information log web requests sent
type ClientRequestLog struct {
	URI    string
	Header map[string][]string
	Errors []error
}

// WriteInfo logs the web request
func (cl ClientRequestLog) WriteInfo(w io.Writer) error {
	log := map[string]interface{}{
		"time":  time.Now().Format(time.RFC3339),
		"type":  "request",
		"level": "info",
		"request": map[string]interface{}{
			"uri":    cl.URI,
			"header": cl.Header,
		},
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

type ClientResponseLog struct {
	Request ClientRequestLog
	Header  map[string][]string
	Status  string
	Errors  []error
}

func (cl ClientResponseLog) WriteInfo(w io.Writer) error {
	log := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"type":    "response",
		"level":   "info",
		"request": cl.Request,
		"response": map[string]interface{}{
			"status": cl.Status,
			"header": cl.Header,
		},
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

func (cl ClientResponseLog) WriteError(w io.Writer) error {
	log := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"type":    "response",
		"level":   "info",
		"request": cl.Request,
		"response": map[string]interface{}{
			"status": cl.Status,
			"header": cl.Header,
		},
		"errors": cl.Errors,
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

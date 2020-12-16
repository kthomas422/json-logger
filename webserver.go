package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// WebServerRequest logs a request to a web server
type ServerRequestLog struct {
	Origin string              // ip of where the request came from
	URI    string              // full uri of request
	Header map[string][]string // header values
	Body   interface{}         // not sure what to expect yet
	Errors []error             // only used for when an error occurs with the request
}

func (sl ServerRequestLog) WriteInfo(w io.Writer) error {
	log := map[string]interface{}{
		"time":  time.Now().Format(time.RFC3339),
		"type":  "request",
		"level": "info",
		"request": map[string]interface{}{
			"uri":    sl.URI,
			"origin": sl.Origin,
			"header": sl.Header,
			"body":   sl.Body,
		},
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

func (sl ServerRequestLog) WriteError(w io.Writer) error {
	log := map[string]interface{}{
		"time":  time.Now().Format(time.RFC3339),
		"type":  "request",
		"level": "error",
		"request": map[string]interface{}{
			"uri":    sl.URI,
			"origin": sl.Origin,
			"header": sl.Header,
			"body":   sl.Body,
		},
		"errors": sl.Errors,
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

// ServerResponseLog logs the response the web server sends
type ServerResponseLog struct {
	Request ServerRequestLog
	Header  map[string]string
	Status  string
	Body    interface{}
	Errors  []error
}

func (sl ServerResponseLog) WriteInfo(w io.Writer) error {
	log := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"type":    "response",
		"level":   "info",
		"request": sl.Request,
		"response": map[string]interface{}{
			"status": sl.Status,
			"header": sl.Header,
			"body":   sl.Body,
		},
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

func (sl ServerResponseLog) WriteError(w io.Writer) error {
	log := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"type":    "response",
		"level":   "error",
		"request": sl.Request,
		"response": map[string]interface{}{
			"status": sl.Status,
			"header": sl.Header,
			"body":   sl.Body,
		},
		"errors": sl.Errors,
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

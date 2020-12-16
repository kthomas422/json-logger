package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type ClientRequestLog struct {
	URI    string
	Header map[string][]string
	Body   interface{}
	Errors []error
}

func (cl ClientRequestLog) WriteInfo(w io.Writer) error {
	log := map[string]interface{}{
		"time":  time.Now().Format(time.RFC3339),
		"type":  "request",
		"level": "info",
		"request": map[string]interface{}{
			"uri":    cl.URI,
			"header": cl.Header,
			"body":   cl.Body,
		},
	}
	logBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(logBytes))
	return err
}

func (cl ClientRequestLog) WriteError(w io.Writer) error {
	log := map[string]interface{}{
		"time":  time.Now().Format(time.RFC3339),
		"type":  "request",
		"level": "error",
		"request": map[string]interface{}{
			"uri":    cl.URI,
			"header": cl.Header,
			"body":   cl.Body,
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

type ClientResponseLog struct {
	Request ClientRequestLog
	Header  map[string][]string
	Status  string
	Body    interface{}
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
			"body":   cl.Body,
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
			"body":   cl.Body,
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

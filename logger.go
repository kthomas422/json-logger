package jsonlogger

import (
	"io"
)

// Logger is the object that stores the output
type Logger struct {
	output io.Writer
}

// InfoLog is an interface for writing info messages to the output.
// See simple.go for an example.
type InfoLog interface {
	WriteInfo(io.Writer) error
}

// ErrorLog is an interface for writing error messages to the output.
// See simple.go for an example.
type ErrLog interface {
	WriteError(io.Writer) error
}

// NewLogger initializes a logger with the output
func NewLogger(output io.Writer) *Logger {
	if output == nil {
		return nil
	}
	return &Logger{
		output: output,
	}
}

// Info writes an "informational" message to the log output
func (l *Logger) Info(lg InfoLog) error {
	return lg.WriteInfo(l.output)
}

// Error writes an "error" message to the log output
func (l *Logger) Error(lg ErrLog) error {
	return lg.WriteError(l.output)
}

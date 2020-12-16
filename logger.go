package logger

import (
	"io"
)

// Logger is the object that stores the output
type Logger struct {
	output io.Writer
}

// Log is an interface for writing info or error messages to the output.
// See the simple.go for an example.
type Log interface {
	WriteInfo(io.Writer) error
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
func (l *Logger) Info(lg Log) error {
	return lg.WriteInfo(l.output)
}

// Error writes an "error" message to the log output
func (l *Logger) Error(lg Log) error {
	return lg.WriteError(l.output)
}

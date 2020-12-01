package logger

import (
	"io"
)

// Logger is the object that stores the output and whether the debugging flag
// is set or not
type Logger struct {
	output io.Writer
	debug  bool
}

// Log is an interface for writing info, debug or error messages to the output.
// See the simple.go for an example.
type Log interface {
	WriteInfo(io.Writer) error
	WriteDebug(io.Writer) error
	WriteError(io.Writer) error
}

// New initializes a logger with the output and the debugging flag
func New(debug bool, output io.Writer) *Logger {
	if output == nil {
		return nil
	}
	return &Logger{
		debug:  debug,
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

// Debug writes a "debugging" message to the log output if the debugging
// flag is set.
func (l *Logger) Debug(lg Log) error {
	if l.debug {
		return lg.WriteDebug(l.output)
	}
	return nil
}

package stacktrace

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime/debug"
	"strings"
)

var re = regexp.MustCompile(`\((0x.+|\.\.\.)\)`)

// wrapper for including stacktrace alongside errors
type wrapper struct {
	err   error
	stack []byte
}

// Error returns the original error string followed by runtime stacktrace
func (w *wrapper) Error() string {
	return fmt.Sprintf("%v\n%s", w.err, format(string(w.stack)))
}

func (w *wrapper) Unwrap() error {
	return w.err
}

// Wrap wraps an error including runtime stacktrace
func Wrap(err error) error {
	if _, ok := err.(*wrapper); ok {
		return err
	}
	return &wrapper{
		err:   err,
		stack: debug.Stack(),
	}
}

func format(stack string) string {
	var buffer bytes.Buffer
	lines := strings.Split(stack, "\n")

	if len(lines) <= 5 {
		return stack
	}

	lines = lines[1:]
	for i := 5; i < len(lines); i += 2 {
		file := strings.Split(lines[i], " ")[0]
		function := re.ReplaceAllString(lines[i-1], "")

		_, err := buffer.WriteString(
			fmt.Sprintf("-- %s -- %s\n", strings.TrimSpace(file), function),
		)
		if err != nil {
			return stack
		}
	}

	return buffer.String()
}

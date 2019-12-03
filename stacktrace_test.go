package stacktrace_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/indebted-modules/stacktrace"
	"github.com/stretchr/testify/suite"
)

type StackTraceSuite struct {
	suite.Suite
}

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func TestStackTraceSuite(t *testing.T) {
	suite.Run(t, new(StackTraceSuite))
}

func (s *StackTraceSuite) TestWrap() {
	err := fmt.Errorf("")
	wrapped := stacktrace.Wrap(err)
	s.Error(wrapped)
	s.True(len(err.Error()) == 0)
	s.True(len(wrapped.Error()) > 0)
}

func (s *StackTraceSuite) TestAs() {
	s.True(errors.As(stacktrace.Wrap(&CustomError{Message: "Too bad!"}), new(*CustomError)))
	s.False(errors.As(stacktrace.Wrap(fmt.Errorf("another error")), new(*CustomError)))
}

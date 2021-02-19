package transport_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ThomasVNN/go-base/log"
	"github.com/ThomasVNN/go-base/transport"
)

func TestLogErrorHandler(t *testing.T) {
	var output []interface{}

	logger := log.Logger(log.LoggerFunc(func(keyvals ...interface{}) error {
		output = append(output, keyvals...)
		return nil
	}))

	errorHandler := transport.NewLogErrorHandler(logger)

	err := errors.New("error")

	errorHandler.Handle(context.Background(), err)

	if output[1] != err {
		t.Errorf("expected an error log event: have %v, want %v", output[1], err)
	}
}

package rate_limiter

import (
	"github.com/pkg/errors"
	"testing"
	"time"
)

func TestThrowsErrorOnTimeout(t *testing.T) {
	setUp(t)

	l := New(1, 1*time.Millisecond)
	errChan := make(chan error)

	go func() {
		_, err := l.Run(func() (interface{}, error) {
			time.Sleep(2 * time.Millisecond)
			return nil, nil
		})
		errChan <- err
	}()

	go func() {
		_, err := l.Run(func() (interface{}, error) {
			time.Sleep(2 * time.Millisecond)
			return nil, nil
		})
		errChan <- err
	}()

	err1 := <-errChan
	err2 := <-errChan

	if err1 != nil && err2 != nil {
		t.Fatalf("Both results are errors <%v>, <%v>", err1, err2)
	}

	if err1 == nil && err2 == nil {
		t.Fatalf("Both results are not errors <%v>, <%v>", err1, err2)
	}

	if err1 != nil && !errors.Is(err1, ErrTimeout) {
		t.Fatalf("Wrong type of error, <%v>", err1)
	}

	if err2 != nil && !errors.Is(err2, ErrTimeout) {
		t.Fatalf("Wrong type of error, <%v>", err2)
	}
}

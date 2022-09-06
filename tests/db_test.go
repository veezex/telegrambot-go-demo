//go:build integration
// +build integration

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/veezex/homework/tests/fixture"
	"testing"
)

func TestDbAppleAdd(t *testing.T) {
	t.Run("can be executed", func(t *testing.T) {
		storage, tearDown := setUpDb(t)
		defer tearDown()

		apple := fixture.Apple().
			ColorName("yellow").
			Price(100).V()

		err := storage.Add(context.Background(), &apple)
		assert.NoError(t, err)
	})
}

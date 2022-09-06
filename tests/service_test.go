//go:build integration
// +build integration

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/tests/fixture"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestServiceAppleGet(t *testing.T) {
	t.Run("can be executed", func(t *testing.T) {
		client, tearDown := setUpClient(t)
		defer tearDown()

		apple := fixture.Apple().
			ColorName("yellow").
			Price(100).V()

		err := client.Add(context.Background(), &apple)
		assert.NoError(t, err)

		// check if it was added
		resultApple, err := client.Get(context.Background(), apple.Id)
		assert.Equal(t, apple, *resultApple)
	})
}

func TestServiceAppleDelete(t *testing.T) {
	t.Run("can be executed", func(t *testing.T) {
		client, tearDown := setUpClient(t)
		defer tearDown()

		apple := fixture.Apple().
			ColorName("yellow").
			Price(100).V()

		// add the apple
		err := client.Add(context.Background(), &apple)
		require.NoError(t, err)

		// delete the apple
		err = client.Delete(context.Background(), apple.Id)
		require.NoError(t, err)

		// check if it was deleted
		_, err = client.Get(context.Background(), apple.Id)
		code, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, code.Code(), codes.NotFound)
	})
}

func TestServiceAppleList(t *testing.T) {
	t.Run("can be executed", func(t *testing.T) {
		client, tearDown := setUpClient(t)
		defer tearDown()

		apple := fixture.Apple().
			ColorName("yellow").
			Price(100).V()

		// add the apple
		err := client.Add(context.Background(), &apple)
		require.NoError(t, err)

		// get the list
		apples, err := client.List(context.Background(), nil)
		require.NoError(t, err)

		assert.Equal(t, apples, []applePkg.Apple{apple})
	})
}

func TestServiceAppleUpdate(t *testing.T) {
	t.Run("can be executed", func(t *testing.T) {
		client, tearDown := setUpClient(t)
		defer tearDown()

		apple := fixture.Apple().
			ColorName("yellow").
			Price(100).V()

		// add the apple
		err := client.Add(context.Background(), &apple)
		require.NoError(t, err)

		// update the apple
		newApple := apple
		newApple.Color.Name = "red"
		newApple.Price = 200
		err = client.Update(context.Background(), &newApple)
		require.NoError(t, err)

		// check if it was deleted
		resultApple, err := client.Get(context.Background(), apple.Id)
		require.NoError(t, err)

		assert.Equal(t, *resultApple, newApple)
	})
}

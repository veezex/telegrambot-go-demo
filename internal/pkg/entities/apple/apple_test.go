package apple

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities"
	"testing"
)

func TestApple(t *testing.T) {
	t.Run("can be stringified", func(t *testing.T) {
		f := setUp(t)

		assert.Equal(t, f.apple.String(), fmt.Sprintf("Apple %d: color %s / price %f", f.apple.Id, f.apple.Color.String(), f.apple.Price))
	})

	t.Run("correct apple should be validated", func(t *testing.T) {
		f := setUp(t)

		require.NoError(t, Validate(&f.apple))
	})

	t.Run("should throw validation error if price below 0", func(t *testing.T) {
		f := setUp(t)
		f.apple.Price = -1

		assert.ErrorIs(t, Validate(&f.apple), entities.ErrValidation)
	})
}

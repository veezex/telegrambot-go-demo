package color

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities"
	"testing"
)

func TestColor(t *testing.T) {
	t.Run("can be stringified", func(t *testing.T) {
		f := setUp(t)

		assert.Equal(t, f.color.String(), fmt.Sprintf("Color %d: color %s", f.color.Id, f.color.Name))
	})

	t.Run("correct color should be validated", func(t *testing.T) {
		f := setUp(t)

		require.NoError(t, Validate(&f.color))
	})

	t.Run("should throw validation error color name is more than 100 characters", func(t *testing.T) {
		f := setUp(t)
		f.color.Name = fake().LongColor

		assert.ErrorIs(t, Validate(&f.color), entities.ErrValidation)
	})
}

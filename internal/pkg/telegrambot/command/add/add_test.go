package add

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	colorPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/color"
	cmdPkg "gitlab.ozon.dev/veezex/homework/internal/pkg/telegrambot/command"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("can add an apple", func(t *testing.T) {
		f := setUp(t)
		price := fake().Price
		color := fake().Color
		a := apple.Apple{
			Id: 0,
			Color: colorPkg.Color{
				Id:   0,
				Name: color,
			},
			Price: float64(price),
		}

		f.storage.EXPECT().
			Add(gomock.Any(), &a).
			Return(nil).
			Times(1)

		s, err := f.impl.Execute(fmt.Sprintf("%s %d", color, price))

		require.NoError(t, err)
		assert.Equal(t, s, a.String()+" was added")
	})

	t.Run("throws an error, when arguments count is not right", func(t *testing.T) {
		f := setUp(t)
		_, err := f.impl.Execute("yellow 1000 red")

		assert.ErrorIs(t, err, cmdPkg.ErrBadArgument)
	})

	t.Run("throws an error, when price is not a number", func(t *testing.T) {
		f := setUp(t)
		_, err := f.impl.Execute("yellow pricestring")

		assert.ErrorIs(t, err, cmdPkg.ErrBadArgument)
	})
}

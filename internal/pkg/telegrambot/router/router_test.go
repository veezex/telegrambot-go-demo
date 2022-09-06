package router

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRouter(t *testing.T) {
	t.Run("can get router summary string", func(t *testing.T) {
		f := setUp(t)
		name := fake().CommandName
		description := fake().CommandDesc

		f.command.EXPECT().
			Name().Return(name).Times(1)

		err := f.router.RegisterCommand(f.command)

		f.command.EXPECT().
			Name().Return(name).Times(1)
		f.command.EXPECT().
			Description().Return(description).Times(1)

		require.NoError(t, err)
		assert.Equal(t, f.router.String(), fmt.Sprintf("/%s - %s", name, description))
	})

	t.Run("can handle command", func(t *testing.T) {
		f := setUp(t)
		name := fake().CommandName

		f.command.EXPECT().
			Name().Return(name).Times(1)

		err := f.router.RegisterCommand(f.command)
		require.NoError(t, err)

		args := fake().CommandArgs
		result := fake().CommandResult
		f.command.EXPECT().
			Execute(args).Return(result, nil).Times(1)

		str, err := f.router.HandleCommand(name, args)
		require.NoError(t, err)
		assert.Equal(t, str, result)
	})

	t.Run("throws an error when handling unknown command", func(t *testing.T) {
		f := setUp(t)
		_, err := f.router.HandleCommand(fake().CommandName, "")

		assert.ErrorIs(t, err, errCommandUnknown)
	})

}

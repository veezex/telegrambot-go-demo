package server

import (
	"context"
	"github.com/pkg/errors"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

func addCmd(ctx context.Context, storage storage.AppleMutableStorage, data interface{}) error {
	appleMap, ok := data.(map[string]interface{})
	if !ok {
		return errors.Wrapf(errUnknownPayload, "add: it should be a map <%v>", data)
	}

	apple, err := applePkg.ParseFromMap(appleMap)
	if err != nil {
		return err
	}

	return storage.Add(ctx, apple)
}

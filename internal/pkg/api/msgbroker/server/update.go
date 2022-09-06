package server

import (
	"context"
	"github.com/pkg/errors"
	applePkg "gitlab.ozon.dev/veezex/homework/internal/pkg/entities/apple"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

func updateCmd(ctx context.Context, storage storage.AppleMutableStorage, data interface{}) error {
	appleMap, ok := data.(map[string]interface{})
	if !ok {
		return errors.Wrapf(errUnknownPayload, "update: it should be a map <%v>", data)
	}

	apple, err := applePkg.ParseFromMap(appleMap)
	if err != nil {
		return err
	}

	return storage.Update(ctx, apple)
}

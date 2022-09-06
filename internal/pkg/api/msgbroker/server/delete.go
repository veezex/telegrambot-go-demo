package server

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/veezex/homework/internal/pkg/storage"
)

func deleteCmd(ctx context.Context, storage storage.AppleMutableStorage, data interface{}) error {
	id, ok := data.(float64)
	if !ok {
		return errors.Wrapf(errUnknownPayload, "add: it a number <%v>", data)
	}

	return storage.Delete(ctx, uint64(id))
}

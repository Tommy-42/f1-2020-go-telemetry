package repository

import (
	"bytes"
	"context"
)

type Repository interface {
	Store(ctx context.Context, body *bytes.Reader) error
}

package store

import (
	"context"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/vibe"
)

type Store interface {
	CreateVibe(ctx context.Context, scream *vibe.Vibe) error
	GetVibe(ctx context.Context, id string) (*vibe.Vibe, error)
	ListVibes(ctx context.Context) ([]*vibe.Vibe, error)
	DeleteVibe(ctx context.Context, id string) error
}

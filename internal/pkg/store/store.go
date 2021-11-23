package store

import (
	"context"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/scream"
)

type Store interface {
	CreateScream(ctx context.Context, scream *scream.Scream) error
	GetScream(ctx context.Context, id string) (*scream.Scream, error)
	ListScreams(ctx context.Context) ([]*scream.Scream, error)
	DeleteScream(ctx context.Context, id string) error
}

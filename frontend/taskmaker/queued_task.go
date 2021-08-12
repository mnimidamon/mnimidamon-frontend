package taskmaker

import "context"

type QueuedTask interface {
	Label() string
	Execute(ctx context.Context, progress *uint) error
}

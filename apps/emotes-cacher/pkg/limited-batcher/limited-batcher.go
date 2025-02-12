package limitedbatcher

import (
	"context"
	"slices"

	"go.uber.org/ratelimit"
)

type LimitedBatcher[T any] struct {
	limiter   ratelimit.Limiter
	batchSize int
}

func New[T any](batchRate, batchSize int) LimitedBatcher[T] {
	return LimitedBatcher[T]{
		limiter:   ratelimit.New(batchRate),
		batchSize: batchSize,
	}
}

// Batch splits elements into chunks with a predefined size (batch size) and executes a batch
// function on each chunk according to the rate-limit quota.
func (lb *LimitedBatcher[T]) Batch(
	ctx context.Context,
	elements []T,
	batch func(context.Context, []T) error,
) error {
	chunks := slices.Chunk(elements, lb.batchSize)

	for chunk := range chunks {
		lb.limiter.Take()

		if err := batch(ctx, chunk); err != nil {
			return err
		}
	}

	return nil
}

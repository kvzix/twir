package batchlimiter

import (
	"context"

	"go.uber.org/ratelimit"
)

type BatchLimiter[T any] struct {
	limiter   ratelimit.Limiter
	batchSize int
}

func New[T any](batchRate, batchSize int) BatchLimiter[T] {
	limiter := ratelimit.New(batchRate)

	return BatchLimiter[T]{
		limiter:   limiter,
		batchSize: batchSize,
	}
}

// Batched executes batches with provided batch function on elements with pre-configured size and
// rate (which is used for rate-limiting of batches in context of BatchLimiter instance).
func (l *BatchLimiter[T]) Batched(
	ctx context.Context,
	elements []T,
	batch func(context.Context, []T) error,
) error {
	var (
		capacity int
		buffer   = make([]T, l.batchSize)
	)

	for _, element := range elements {
		buffer[capacity] = element
		capacity += 1

		// Batch buffer is full, execute batch on the elements from this buffer.
		if capacity == l.batchSize {
			err := l.limit(func() error {
				return batch(ctx, buffer)
			})
			if err != nil {
				return err
			}

			capacity = 0
		}
	}

	// Make sure that remaining elements are batched in case the buffer was not full
	// at the time of the last batch.
	if capacity != 0 {
		return l.limit(func() error {
			return batch(ctx, buffer[:capacity])
		})
	}

	return nil
}

// limit limits execution of the provided function with rate-limiter.
func (l *BatchLimiter[T]) limit(fn func() error) error {
	l.limiter.Take()
	return fn()
}

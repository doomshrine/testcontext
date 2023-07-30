package testcontext

import (
	"context"
	"testing"
	"time"
)

var DefaultTimeout = time.Second * 30

func FromT(ctx context.Context, t *testing.T) (context.Context, context.CancelFunc) {
	return FromTimeout(ctx, t, DefaultTimeout)
}

func FromTimeout(ctx context.Context, t *testing.T, timeout time.Duration) (context.Context, context.CancelFunc) {
	d := time.Now().Add(timeout)
	return FromDeadline(ctx, t, d)
}

func FromDeadline(ctx context.Context, t *testing.T, deadline time.Time) (context.Context, context.CancelFunc) {
	testDeadline, ok := t.Deadline()
	if ok {
		return fromTimes(ctx, testDeadline, deadline)
	}

	return fromTimes(ctx, testDeadline, testDeadline)
}

func fromTimes(ctx context.Context, t1, t2 time.Time) (context.Context, context.CancelFunc) {
	if t1.Before(t2) {
		return context.WithDeadline(ctx, t1)
	}

	return context.WithDeadline(ctx, t2)
}

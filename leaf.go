package leaf

import (
	"context"
	"time"

	"github.com/isaqueveras/leaf/internal"
)

// New initialize a new queue
func New(ctx context.Context, workers, itemsPerPage uint64, interval time.Duration) internal.IQueue {
	return internal.New(ctx, workers, itemsPerPage, interval)
}

// GetPage get current page data
func GetPage(ctx context.Context) internal.IPage {
	return internal.GetPage(ctx)
}

// GetData get the data to be processed
func GetData(ctx context.Context) interface{} {
	return internal.GetData(ctx)
}

// Stop stops the execution of all routines
func Stop(ctx context.Context) {
	internal.Stop(ctx)
}

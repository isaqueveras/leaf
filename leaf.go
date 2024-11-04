package leaf

import (
	"context"
	"time"

	"github.com/isaqueveras/leaf/internal"
)

// IQueue defines the interface for the methods of a queue
type IQueue internal.IQueue

// New initialize a new queue
func New(ctx context.Context, workers, itemsPerPage uint64, interval time.Duration) IQueue {
	return internal.New(ctx, workers, itemsPerPage, interval)
}

// IPage defines the methods of a page
type IPage internal.IPage

// GetPage get current page data
func GetPage(ctx context.Context) IPage {
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

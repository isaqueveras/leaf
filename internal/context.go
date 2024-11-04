package internal

import "context"

type contextKey int

const (
	pageKey contextKey = iota + 1
	dataKey
	stopKey
)

// GetPage get current page data
func GetPage(ctx context.Context) IPage {
	if value, ok := ctx.Value(pageKey).(page); ok {
		return &value
	}
	return nil
}

// GetData get the data to be processed
func GetData(ctx context.Context) interface{} {
	return ctx.Value(dataKey)
}

// Stop stops the execution of all routines
func Stop(ctx context.Context) {
	if stop, ok := ctx.Value(stopKey).(func()); ok {
		stop()
	}
}

func contextWithPage(parent context.Context, page page) context.Context {
	return context.WithValue(parent, pageKey, page)
}

func contextWithData(parent context.Context, data interface{}) context.Context {
	return context.WithValue(parent, dataKey, data)
}

func contextWithStop(parent context.Context, stop func()) context.Context {
	return context.WithValue(parent, stopKey, stop)
}

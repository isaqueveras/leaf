package internal

import (
	"context"
	"time"
)

// IQueue defines the interface for the methods of a queue
type IQueue interface {
	// Publish defines the method for publishing data to the queue
	Publish(PublishFunc) *queue
	// Consume defines the method to consume data from the queue
	Consume(ConsumerFunc) *queue
	// Wait defines the method to wait for publishers and consumers to execute
	Wait()
}

// ConsumerFunc defines the type of function to consume the data
type ConsumerFunc func(context.Context) error

// Consume implements the method to consume the data
func (fn ConsumerFunc) Consume(ctx context.Context) error {
	return fn(ctx)
}

// PublishFunc defines the type of function to publish the data
type PublishFunc func(context.Context) (interface{}, error)

// Publish implements the method to publish the data
func (fn PublishFunc) Publish(ctx context.Context) (interface{}, error) {
	return fn(ctx)
}

type queue struct {
	interval time.Duration
	pipe     chan interface{}
	group    iGroup
	pool     iWorkerPool
	page     IPage
	quit     chan bool
	ctx      context.Context
}

// New initialize a new queue
func New(ctx context.Context, workers, itemsPerPage uint64, interval time.Duration) IQueue {
	if workers < 3 {
		workers += 3
	}

	return &queue{
		ctx:      ctx,
		interval: interval,
		group:    newGroup(ctx),
		page:     newPage(itemsPerPage),
		pool:     newWorkerPool(workers),
		quit:     make(chan bool, 1),
		pipe:     make(chan interface{}),
	}
}

// Wait implements the method to wait for the queue to execute
func (queue *queue) Wait() {
	queue.group.wait()
	<-queue.quit
}

func (queue *queue) stop() {
	queue.quit <- true
	close(queue.pipe)
	queue.pool.close()
	queue.group.cancel()
	close(queue.quit)
}

// Consume implements the method to consume the data
func (queue *queue) Consume(fn ConsumerFunc) *queue {
	queue.pool.schedule(queue.group.wrap(consume(fn, queue)))
	return queue
}

func consume(fn ConsumerFunc, queue *queue) func() error {
	return func() error {
		for value := range queue.pipe {
			queue.pool.schedule(queue.group.wrap(func() error {
				return fn.Consume(contextWithStop(contextWithData(queue.ctx, value), queue.stop))
			}))
		}
		return nil
	}
}

// Publish implements the method to publish the data
func (queue *queue) Publish(fn PublishFunc) *queue {
	queue.pool.schedule(queue.group.wrap(publish(fn, queue)))
	return queue
}

func publish(fn PublishFunc, queue *queue) func() error {
	return func() error {
		for range time.NewTicker(queue.interval).C {
			queue.pool.schedule(queue.group.wrap(func() error {
				queue.page.calculate()
				data, err := fn.Publish(contextWithStop(contextWithPage(queue.ctx, queue.page.getPage()), queue.stop))
				if err != nil {
					return err
				}
				queue.pipe <- data
				return nil
			}))
		}
		return nil
	}
}

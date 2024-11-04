package main

import (
	"context"
	"log"
	"time"

	"github.com/isaqueveras/leaf"
)

func main() {
	// 1 - context
	// 2 - total number of workers
	// 3 - total items per page
	// 4 - interval at which the publish function will be executed
	queue := leaf.New(context.Background(), 50, 100, time.Second)

	queue.Publish(publish)
	queue.Consume(consume)

	queue.Wait()
}

func consume(ctx context.Context) error {
	value := leaf.GetData(ctx).(time.Time)
	log.Printf("[consumer] - Value: %v\n", value.String())
	return nil
}

func publish(ctx context.Context) (interface{}, error) {
	page := leaf.GetPage(ctx)

	log.Printf("[publisher] - Page: %d Offset: %d ItemsPerPage %d",
		page.GetCurrentPage(),
		page.GetOffset(),
		page.GetItemsPerPage(),
	)

	if page.GetCurrentPage() == 5000 {
		leaf.Stop(ctx)
	}

	return time.Now(), nil
}

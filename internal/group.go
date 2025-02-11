package internal

import (
	"context"
	"log"
	"sync"
)

type group struct {
	ctx context.Context
	wg  sync.WaitGroup
}

type iGroup interface {
	wrap(func() error) func()
	wait()
}

func newGroup(parent context.Context) iGroup {
	return &group{ctx: parent, wg: sync.WaitGroup{}}
}

// wait for any packaged goroutine to finish and return an error that occurred
func (g *group) wait() {
	log.Println("Waiting on Group")
	g.wg.Wait()
}

// wraps a function to be executed in a goroutine
func (g *group) wrap(fn func() error) func() {
	g.wg.Add(1)
	return func() {
		defer g.wg.Done()

		_, cancel := context.WithCancel(g.ctx)
		defer cancel()

		if err := fn(); err != nil {
			log.Println("Group canceling its context due to error: ", err.Error())
		}
	}
}

package internal

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type group struct {
	wg         *sync.WaitGroup
	ctx        context.Context
	cancelCtx  context.CancelFunc
	errorMutex sync.Mutex
}

type iGroup interface {
	wrap(func() error) func()
	wait()
	cancel()
}

func newGroup(parent context.Context) iGroup {
	return &group{ctx: parent, wg: &sync.WaitGroup{}}
}

func (g *group) cancel() {
	g.cancelCtx()
}

// wait for any packaged goroutine to finish and return an error that occurred
func (g *group) wait() {
	log.Println("Waiting on Group")
	g.wg.Wait()

	if g.cancelCtx != nil {
		g.cancel()
	}
}

// wraps a function to be executed in a goroutine
func (g *group) wrap(fn func() error) func() {
	g.wg.Add(1)

	return func() {
		defer g.wg.Done()
		_, g.cancelCtx = context.WithCancel(g.ctx)

		if err := fn(); err != nil {
			g.errorMutex.Lock()
			fmt.Printf("Error: %v\n", err.Error())
			if g.cancelCtx != nil {
				log.Println("Group canceling its context due to error")
				g.cancel()
			}
			g.errorMutex.Unlock()
		}
	}
}

package internal

import (
	"fmt"
	"log"
)

type scheduleFunc func()

func (fn scheduleFunc) Schedule() { fn() }

type iWorkerPool interface {
	schedule(scheduleFunc)
	close()
}

type workerPool struct {
	jobs chan scheduleFunc
}

// newWorkerPool constructs a new worker pool
func newWorkerPool(total uint64) iWorkerPool {
	log.Printf("Creating worker pool with %d workers", total)

	jobs := make(chan scheduleFunc)
	for i := 0; i < int(total); i++ {
		go worker(jobs)
	}

	return &workerPool{jobs: jobs}
}

func (p *workerPool) schedule(job scheduleFunc) {
	p.jobs <- job
}

func (p *workerPool) close() {
	close(p.jobs)
}

func worker(jobs chan scheduleFunc) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover: %v\n", r)
			return
		}
	}()

	for job := range jobs {
		job.Schedule()
	}
}

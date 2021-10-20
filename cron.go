package kgo

import (
	"sync"
	"time"
)

// CronJob ...
type CronJob struct {
	TimeThreshold time.Duration
	Period        time.Duration
	running       bool
	mutex         sync.RWMutex
	wg            sync.WaitGroup
}

func (cj *CronJob) exec(fx func()) {
	defer cj.wg.Done()
	running := cj.running
	for running {
		t := time.Now()
		fx()
		dt := time.Since(t)
		for running && dt < cj.TimeThreshold {
			time.Sleep(cj.Period)
			dt = time.Since(t)
			cj.mutex.RLock()
			running = cj.running
			cj.mutex.RUnlock()
		}
	}
}

// Start ...
func (cj *CronJob) Start(fx func()) {
	cj.wg.Add(1)
	cj.running = true
	go cj.exec(fx)
}

// Stop ...
func (cj *CronJob) Stop() {
	cj.mutex.Lock()
	cj.running = false
	cj.mutex.Unlock()
	cj.wg.Wait()
}

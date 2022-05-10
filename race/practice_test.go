// Author: LiuLin
// Date: 2022/4/27

package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func (c *Config) T() {
	//fmt.Printf("%v\n", c)
}

func BenchmarkAt(b *testing.B) {
	cfg := &Config{}
	value := atomic.Value{}
	value.Store(cfg)
	b.ResetTimer()
	go func() {
		i := 0
		for {
			i++
			cfg = &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			value.Store(cfg)
		}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				config := value.Load().(*Config)
				config.T()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkRw(b *testing.B) {
	cfg := &Config{}
	var mutex sync.RWMutex
	b.ResetTimer()
	go func() {
		i := 0
		for {
			i++
			mutex.Lock()
			cfg = &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			mutex.Unlock()
		}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				mutex.RLock()
				cfg.T()
				mutex.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

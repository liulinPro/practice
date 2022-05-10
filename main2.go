package main

import "time"

func main() {

}

func WaitUntilTrueOrMaxTries(maxTries int, wait time.Duration, f func() bool) bool {
	for try := 0; try < maxTries; try++ {
		if ok := f(); ok {
			return true
		}
		time.Sleep(wait)
		wait *= 2
	}
	return false
}

func WaitUntilTrueOrTimeout(timeout, wait time.Duration, f func() bool) bool {
	out := time.NewTimer(timeout)
	for {
		select {
		case <-out.C:
			return false
		default:
			if ok := f(); ok {
				return true
			}
			time.Sleep(wait)
		}
	}
}

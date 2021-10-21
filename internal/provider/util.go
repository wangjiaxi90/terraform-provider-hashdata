package provider

import (
	"math/rand"
	"time"
)

func Int32(v int) *int32 {
	t := int32(v)
	return &t
}
func String(v string) *string {
	return &v
}
func Bool(v bool) *bool {
	return &v
}

func simpleRetry(fn func() error) error {
	return retry(MAX_RETRY_TIMES, 10*time.Second, fn)
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}

func isServerBusy(err error) error {
	if err != nil {
		return stop{err}
	}
	return nil
}

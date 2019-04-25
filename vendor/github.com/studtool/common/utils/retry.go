package utils

import (
	"time"
)

type RetryFunc func(n int) error

func WithRetry(f RetryFunc, n int, d time.Duration) (err error) {
	for i := 0; i <= n; i++ {
		if err = f(i); err == nil {
			return nil
		}
		time.Sleep(d)
	}
	return err
}

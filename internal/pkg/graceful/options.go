package graceful

import "time"

type opt struct {
	stopTimeout time.Duration
	waitTime    time.Duration
}

type Option interface {
	apply(opt *opt)
}

type optFunc func(*opt)

func (f optFunc) apply(args *opt) {
	f(args)
}

func WithStopTimeout(t time.Duration) Option {
	return optFunc(func(o *opt) {
		o.stopTimeout = t
	})
}

func WithWaitTime(t time.Duration) Option {
	return optFunc(func(o *opt) {
		o.waitTime = t
	})
}

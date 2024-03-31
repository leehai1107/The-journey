package xhttp

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type clientOptions struct {
	skipLog         bool
	proxyURL        string
	splitLogBody    bool
	splitLogBodyLen int
	timeout         time.Duration
	promCfg         PromConfig
}

type PromConfig struct {
	Enable     bool
	Subsystem  string
	ConstLabel map[string]string
	Register   prometheus.Registerer
}

func NewDefPromConfig() PromConfig {
	return PromConfig{
		Enable:     true,
		Subsystem:  defaultSubsystem,
		ConstLabel: map[string]string{},
		Register:   prometheus.DefaultRegisterer,
	}
}

func NewBasePromConfig(subSystem, recipient string) PromConfig {
	return PromConfig{
		Enable:    true,
		Subsystem: subSystem,
		ConstLabel: map[string]string{
			"recipient": recipient,
		},
		Register: prometheus.DefaultRegisterer,
	}
}

type Option interface {
	apply(*clientOptions)
}

type optionFunc func(*clientOptions)

func (f optionFunc) apply(args *clientOptions) {
	f(args)
}

func WithSkipLog(skipLog bool) Option {
	return optionFunc(func(args *clientOptions) {
		args.skipLog = skipLog
	})
}

func WithProxyURL(proxyURL string) Option {
	return optionFunc(func(args *clientOptions) {
		args.proxyURL = proxyURL
	})
}

func WithSplitLogBody(splitLen ...int) Option {
	return optionFunc(func(args *clientOptions) {
		args.splitLogBody = true
		if len(splitLen) != 0 {
			args.splitLogBodyLen = splitLen[0]
		}
	})
}

func WithTimeout(timeout time.Duration) Option {
	return optionFunc(func(args *clientOptions) {
		args.timeout = timeout
	})
}

func WithProm() Option {
	return optionFunc(func(args *clientOptions) {
		args.promCfg = NewDefPromConfig()
	})
}

func WithBaseProm(subSystem, recipient string) Option {
	return optionFunc(func(args *clientOptions) {
		args.promCfg = NewBasePromConfig(subSystem, recipient)
	})
}

func WithPromConfig(config PromConfig) Option {
	return optionFunc(func(args *clientOptions) {
		args.promCfg = config
	})
}

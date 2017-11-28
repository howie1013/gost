package gost

import "net"

type CallbackOptions struct {
	OnListened       func(net.Listener, error)
	OnListenerClosed func(net.Listener)
}

type CallbackOption func(opts *CallbackOptions)

func OnListened(cb func(net.Listener, error)) CallbackOption {
	return func(opts *CallbackOptions) {
		opts.OnListened = cb
	}
}

func OnListenerClosed(cb func(net.Listener)) CallbackOption {
	return func(opts *CallbackOptions) {
		opts.OnListenerClosed = cb
	}
}

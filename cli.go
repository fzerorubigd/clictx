package clictx

import (
	"context"
	"os"
	"os/signal"
	"sort"
	"strings"
	"sync"
)

var (
	lock = sync.Mutex{}
	old  = map[string]context.Context{}
)

// Context returns a context that is cancelled automatically when a signal received
// If no signals are provided, all incoming signals cancel the context.
// Otherwise, just the provided signals will.
// The idea is to create one context based on a combination of signalls, and
// not one context for each call.
func Context(signals ...os.Signal) context.Context {
	lock.Lock()
	defer lock.Unlock()

	key := signalKey(signals)
	if ctx, ok := old[key]; ok {
		return ctx
	}

	var sig = make(chan os.Signal, len(signals))
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(sig, signals...)
	go func() {
		<-sig
		cancel()

		// Make sure we delete the key, so the next time -IF- someone create
		// the same watch, they get a new context this time (not a canceled context)
		lock.Lock()
		defer lock.Unlock()

		delete(old, key)
	}()

	old[key] = ctx
	return ctx
}

func signalKey(signals []os.Signal) string {
	sort.Slice(signals, func(i, j int) bool {
		return strings.Compare(signals[i].String(), signals[j].String()) < 0
	})

	var key string
	dup := make(map[string]bool, len(signals))
	for i := 0; i < len(signals); i++ {
		if !dup[signals[i].String()] {
			dup[signals[i].String()] = true
			key += signals[i].String()
		}
	}

	return key
}

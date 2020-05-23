package clictx

import (
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCliContext(t *testing.T) {
	ctx := Context(syscall.SIGUSR2, syscall.SIGUSR1)
	select {
	case <-ctx.Done():
		require.True(t, false, "context canceled early")
	default:
	}
	_ = syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)

	select {
	case <-ctx.Done():
	case <-time.After(time.Second):
		require.True(t, false, "context not canceled")
	}

	// Retry with new context
	ctx = Context(syscall.SIGUSR2, syscall.SIGUSR1)
	select {
	case <-ctx.Done():
		require.True(t, false, "context canceled early")
	default:
	}

	// Get the context again should return the same context
	ctx = Context(syscall.SIGUSR2, syscall.SIGUSR1)
	_ = syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)

	select {
	case <-ctx.Done():
	case <-time.After(time.Second):
		require.True(t, false, "context not canceled")
	}
}

type sig string

func (s sig) String() string {
	return string(s)
}

func (s sig) Signal() {

}

func TestSignalKey(t *testing.T) {
	table := map[string][]os.Signal{
		"": {},
		"1": {
			sig("1"), sig("1"),
		},
		"123": {
			sig("3"), sig("2"), sig("3"), sig("2"), sig("1"),
		},
	}

	for i := range table {
		require.Equal(t, i, signalKey(table[i]))
	}
}

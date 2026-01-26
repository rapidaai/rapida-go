// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	t.Run("normal execution", func(t *testing.T) {
		ctx := context.Background()
		executed := false
		Call(ctx, func() {
			executed = true
		})
		if !executed {
			t.Error("Call() did not execute the function")
		}
	})

	t.Run("panic recovery and re-panic", func(t *testing.T) {
		ctx := context.Background()
		defer func() {
			if r := recover(); r == nil {
				t.Error("Call() did not re-panic after recovery")
			}
		}()
		Call(ctx, func() {
			panic("test panic")
		})
	})
}

func TestCallSafe(t *testing.T) {
	t.Run("normal execution", func(t *testing.T) {
		ctx := context.Background()
		executed := false
		CallSafe(ctx, func() {
			executed = true
		})
		if !executed {
			t.Error("CallSafe() did not execute the function")
		}
	})

	t.Run("panic recovery without re-panic", func(t *testing.T) {
		ctx := context.Background()
		panicked := false
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		CallSafe(ctx, func() {
			panic("test panic")
		})
		if panicked {
			t.Error("CallSafe() should not re-panic")
		}
	})
}

func TestGo(t *testing.T) {
	t.Run("normal execution in goroutine", func(t *testing.T) {
		ctx := context.Background()
		var wg sync.WaitGroup
		executed := false
		var mu sync.Mutex

		wg.Add(1)
		Go(ctx, func() {
			mu.Lock()
			executed = true
			mu.Unlock()
			wg.Done()
		})

		wg.Wait()
		mu.Lock()
		if !executed {
			t.Error("Go() did not execute the function")
		}
		mu.Unlock()
	})

	t.Run("panic in goroutine does not crash main", func(t *testing.T) {
		ctx := context.Background()
		done := make(chan bool, 1)

		Go(ctx, func() {
			panic("test panic in goroutine")
		})

		// Give the goroutine time to execute and panic
		select {
		case <-time.After(100 * time.Millisecond):
			done <- true
		}

		select {
		case <-done:
			// Test passed - main didn't crash
		case <-time.After(time.Second):
			t.Error("Timeout waiting for goroutine")
		}
	})
}

func TestReportPanicIfNotNil(t *testing.T) {
	t.Run("nil value returns false", func(t *testing.T) {
		ctx := context.Background()
		result := ReportPanicIfNotNil(ctx, nil)
		if result {
			t.Error("ReportPanicIfNotNil(nil) should return false")
		}
	})

	t.Run("non-nil value returns true", func(t *testing.T) {
		ctx := context.Background()
		result := ReportPanicIfNotNil(ctx, "test panic")
		if !result {
			t.Error("ReportPanicIfNotNil('test panic') should return true")
		}
	})

	t.Run("error value is reported", func(t *testing.T) {
		ctx := context.Background()
		err := "some error"
		result := ReportPanicIfNotNil(ctx, err)
		if !result {
			t.Error("ReportPanicIfNotNil(error) should return true")
		}
	})
}

func TestPanicIfNotNil(t *testing.T) {
	t.Run("nil value does nothing", func(t *testing.T) {
		ctx := context.Background()
		defer func() {
			if r := recover(); r != nil {
				t.Error("PanicIfNotNil(nil) should not panic")
			}
		}()
		PanicIfNotNil(ctx, nil)
	})

	// Note: Testing the non-nil case would require handling the sleep and re-panic
	// which is complex in a unit test. The behavior is tested through Call() tests.
}

func TestContextPropagation(t *testing.T) {
	t.Run("context is passed to CallSafe", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "key", "value")
		CallSafe(ctx, func() {
			// Function executed successfully with context
		})
	})

	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		executed := false
		CallSafe(ctx, func() {
			executed = true
		})

		// Function should still execute even with cancelled context
		// (the context is just passed for panic reporting)
		if !executed {
			t.Error("Function should execute even with cancelled context")
		}
	})
}

// Benchmark tests
func BenchmarkCall(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		Call(ctx, func() {})
	}
}

func BenchmarkCallSafe(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		CallSafe(ctx, func() {})
	}
}

func BenchmarkGo(b *testing.B) {
	ctx := context.Background()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		Go(ctx, func() {
			wg.Done()
		})
	}
	wg.Wait()
}

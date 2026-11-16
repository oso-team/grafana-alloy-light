package tracing

import (
	"context"
	"testing"
)

func TestDefaultTracerStartsAndStops(t *testing.T) {
	tracer, err := New(DefaultOptions)
	if err != nil {
		t.Fatalf("create default tracer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := tracer.Run(ctx); err != nil {
		t.Fatalf("start and stop default tracer: %v", err)
	}
}

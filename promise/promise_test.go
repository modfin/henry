package promise

import (
	"context"
	"testing"
	"time"
)

func TestPromise_Await1(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	p := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello", nil
	})

	s, err := p.Await()
	if err != nil {
		t.Logf("expected no error, got %v", err)
		t.Fail()
	}
	if s != "Hello" {
		t.Logf("expected return Hello, got %v", s)
		t.Fail()
	}
}

func TestPromise_Await2(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	p := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello", nil
	})

	s, err := Await(p)
	if err != nil {
		t.Logf("expected no error, got %v", err)
		t.Fail()
	}
	if s != "Hello" {
		t.Logf("expected return Hello, got %v", s)
		t.Fail()
	}
}

func TestPromise_AwaitAll(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	p1 := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello 1", nil
	})
	p2 := Async(ctx, func(_ context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "Hello 2", nil
	})

	err := AwaitAll(ctx, p1, p2)
	if err != nil {
		t.Logf("expected no error, got %v", err)
		t.Fail()
	}
	if p1.result != "Hello 1" {
		t.Logf("expected return Hello 1, got %v", p1.result)
		t.Fail()
	}
	if p2.result != "Hello 2" {
		t.Logf("expected return Hello 2, got %v", p2.result)
		t.Fail()
	}
}
